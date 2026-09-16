package main

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dake-edu/gopher-shop/internal/logger"
	"github.com/dake-edu/gopher-shop/internal/middleware"
	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/service"
	"github.com/dake-edu/gopher-shop/internal/store"
	"github.com/dake-edu/gopher-shop/internal/version"
	"github.com/dake-edu/gopher-shop/internal/worker"
)

// PageData creates a standard payload for our templates
type PageData struct {
	Title           string
	Year            int
	Books           []models.Book
	Book            models.Book // For details page
	UpdateAvailable bool
	LatestVersion   string
}

var (
	updateMu sync.RWMutex
	// ⚓ ARCHITECTURE UPGRADE: We now use the Service, not the Store directly.
	latestVersion   = version.Current
	updateAvailable = false
)

func main() {
	// 0. SETUP LOGGER
	logger.Setup()

	// 1. UPDATE CHECK (Background)
	go checkForUpdates()

	// 2. WORKER POOL (Background)
	// ⚓ CONCURRENCY: The Conveyor Belt
	// We create a buffered channel. It can hold 100 orders before blocking.
	orderQueue := make(chan worker.Order, 100)

	// Start the "Factory Worker" in the background
	go worker.StartDispatcher(orderQueue)

	// 3. Initialize Layers (Dependency Injection)
	// Store (The Warehouse)
	repo := store.NewInMemoryBookStore()
	// Service (The Brain)
	bookService := service.NewBookService(repo)

	mux := newDemoHandler(bookService, orderQueue)

	// 4. START
	port := "127.0.0.1:8082"
	slog.Info("starting server", "port", port, "url", "http://"+port)

	// Wrap with Middleware
	handler := middleware.RequestLogger(mux)

	server := &http.Server{Addr: port, Handler: handler, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second, IdleTimeout: time.Minute}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server crashed", "error", err)
		os.Exit(1)
	}
}

// render parses the layout files plus the specific page template
func render(w http.ResponseWriter, pageTemplate string, data interface{}) {
	// 1. Define the files to parse
	// We always need base, header, footer, sidebar.
	files := []string{
		"cmd/web-demo/templates/layouts/base.html",
		"cmd/web-demo/templates/parts/head.html",
		"cmd/web-demo/templates/parts/header.html",
		"cmd/web-demo/templates/parts/footer.html",
		"cmd/web-demo/templates/parts/sidebar.html",
		"cmd/web-demo/templates/" + pageTemplate, // e.g., pages/home.html
	}

	// 2. Parse them
	// Note: In a real app, you would cache these templates (parse once, execute many).
	// For this demo, parsing on every request allows you to edit HTML without restarting.
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		slog.Error("template parse error", "error", err, "template", pageTemplate)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 3. Execute "base" (because base.html defines the outline)
	var output bytes.Buffer
	if err := tmpl.ExecuteTemplate(&output, "base", data); err != nil {
		log.Printf("Template Execute Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := output.WriteTo(w); err != nil {
		log.Printf("Template write error: %v", err)
	}
}

// checkForUpdates polls the GitHub repository for the latest version.
func checkForUpdates() {
	// Give the server a moment to start before the first check
	time.Sleep(2 * time.Second)

	// In a real app, you might check once a day.
	// For this demo, we check on startup.
	url := "https://raw.githubusercontent.com/dake-edu/gopher-shop/main/version.txt"

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		slog.Warn("update check failed", "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("update check failed", "status", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 129))
	if err != nil {
		slog.Warn("failed to read version file", "error", err)
		return
	}

	remoteVersion := strings.TrimSpace(string(body))
	if len(body) > 128 || remoteVersion == "" {
		return
	}
	updateMu.Lock()
	defer updateMu.Unlock()
	if remoteVersion != version.Current {
		slog.Info("new version available", "current", version.Current, "remote", remoteVersion)
		updateAvailable = true
		latestVersion = remoteVersion
	} else {
		slog.Info("running latest version", "version", version.Current)
	}
}

func updateStatus() (bool, string) {
	updateMu.RLock()
	defer updateMu.RUnlock()
	return updateAvailable, latestVersion
}

func newDemoHandler(bookService *service.BookService, orderQueue chan<- worker.Order) http.Handler {
	// 4. HANDLERS
	mux := http.NewServeMux()

	// GET / (Home with Filter)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// 🧠 Service handles filtering logic now!
		category := r.URL.Query().Get("category")
		displayBooks, err := bookService.ListBooks(category)
		if err != nil {
			slog.Error("failed to fetch books", "error", err)
			http.Error(w, "Could not fetch books", http.StatusInternalServerError)
			return
		}

		// 3. Render Template
		available, latest := updateStatus()
		data := PageData{
			Title:           "The Gopher Shop",
			Year:            2026,
			Books:           displayBooks,
			UpdateAvailable: available,
			LatestVersion:   latest,
		}

		render(w, "pages/home.html", data)
	})

	// GET /book/{id} (View Details)
	mux.HandleFunc("GET /book/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		book, err := bookService.GetBook(id)
		if err != nil {
			// In a real app, we'd check if err is "NotFound"
			if errors.Is(err, service.ErrBookNotFound) {
				slog.Warn("book not found", "id", id)
				http.NotFound(w, r)
				return
			}
			slog.Error("database error fetching book", "id", id, "error", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		available, latest := updateStatus()
		data := PageData{
			Title:           book.Title,
			Year:            time.Now().Year(),
			Book:            book,
			UpdateAvailable: available,
			LatestVersion:   latest,
		}

		render(w, "pages/details.html", data)
	})

	// GET /add (Show Add Form)
	mux.HandleFunc("GET /add", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Add New Book",
			Year:  time.Now().Year(),
		}
		render(w, "pages/add.html", data)
	})

	// POST /add (Form Submission)
	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		// Bound and validate the form before using its values.
		r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid or oversized form", http.StatusBadRequest)
			return
		}
		price, err := strconv.ParseFloat(r.PostForm.Get("price"), 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}

		newBook := &models.Book{
			Title:    r.FormValue("title"),
			Author:   r.FormValue("author"),
			Price:    price,
			Category: r.FormValue("category"),
			ImageURL: r.FormValue("image_url"),
		}

		// ⚓ VISUAL ANCHOR: The Quality Gate
		// Service Layer handles validation now!
		if err := bookService.CreateBook(newBook); err != nil {
			slog.Warn("book creation failed validation", "title", newBook.Title, "error", err)
			// If validation failed, we should probably show the error to user.
			// For this demo, just 400.
			http.Error(w, "Failed to save book: "+err.Error(), http.StatusBadRequest)
			return
		}

		slog.Info("book created", "title", newBook.Title)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// GET /checkout/{id} (Payment Page)
	mux.HandleFunc("GET /checkout/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		book, err := bookService.GetBook(id)
		if err != nil {
			if errors.Is(err, service.ErrBookNotFound) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title: "Checkout - " + book.Title,
			Year:  time.Now().Year(),
			Book:  book,
		}
		render(w, "pages/checkout.html", data)
	})

	// POST /checkout (Process Payment)
	mux.HandleFunc("POST /checkout", func(w http.ResponseWriter, r *http.Request) {
		// This queues a local demonstration job; no payment is taken.
		r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid or oversized form", http.StatusBadRequest)
			return
		}
		idStr := r.PostForm.Get("book_id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		if _, err := bookService.GetBook(id); err != nil {
			if errors.Is(err, service.ErrBookNotFound) {
				http.NotFound(w, r)
			} else {
				http.Error(w, "Could not load book", http.StatusInternalServerError)
			}
			return
		}

		// 2. Queue a demonstration job
		// Instead of doing work here (Blocking), we send it to the factory (Non-blocking).
		order := worker.Order{
			ID:    id,
			Email: "customer@example.com", // Mock
		}

		// ⚓ CONCURRENCY: Send to Channel
		// A successful send only confirms an in-memory enqueue, not durable storage.
		select {
		case orderQueue <- order:
			slog.Info("order queued", "order_id", id)
		default:
			slog.Error("order queue full", "order_id", id)
			http.Error(w, "System overloaded, try again later", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Demo job queued in memory. No payment was taken or purchase saved."))
	})

	// Serve Static Assets
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("cmd/web-demo/assets"))))

	return mux
}
