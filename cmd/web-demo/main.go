package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/store"
)

// PageData creates a standard payload for our templates
type PageData struct {
	Title string
	Year  int
	Books []models.Book
	Book  models.Book // For details page
}

// Declare tmpl as a package variable isn't needed anymore for this strategy,
// but we might want to cache the "base" parts.
// For simplicity and to avoid race conditions in this learning demo,
// we will parse on every request (development mode style) or use a map.
// Let's use a helper function.

var bookStore *store.InMemoryBookStore

func main() {
	// 1. Initialize Store
	bookStore = store.NewInMemoryBookStore()

	// 2. HANDLERS
	mux := http.NewServeMux()

	// GET / (Home with Filter)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		allBooks, err := bookStore.GetAll()
		if err != nil {
			http.Error(w, "Could not fetch books", http.StatusInternalServerError)
			return
		}

		// FILTERING
		category := r.URL.Query().Get("category")
		var displayBooks []models.Book
		if category != "" {
			for _, b := range allBooks {
				if b.Category == category {
					displayBooks = append(displayBooks, b)
				}
			}
		} else {
			displayBooks = allBooks
		}

		data := PageData{
			Title: "Home",
			Year:  time.Now().Year(),
			Books: displayBooks,
		}

		render(w, "pages/home.html", data)
	})

	// GET /book/{id} (View Details)
	mux.HandleFunc("GET /book/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		var id int
		fmt.Sscanf(idStr, "%d", &id)

		book, found, err := bookStore.GetByID(id)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		if !found {
			http.NotFound(w, r)
			return
		}

		data := PageData{
			Title: book.Title,
			Year:  time.Now().Year(),
			Book:  book,
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
		// Simple Form Parsing
		var price float64
		fmt.Sscanf(r.FormValue("price"), "%f", &price)

		newBook := &models.Book{
			Title:    r.FormValue("title"),
			Author:   r.FormValue("author"),
			Price:    price,
			Category: r.FormValue("category"),
		}

		// ⚓ VISUAL ANCHOR: The Quality Gate
		if newBook.Title == "" {
			http.Error(w, "Title is required (Quality Gate Closed)", http.StatusBadRequest)
			return
		}

		if err := bookStore.Create(newBook); err != nil {
			http.Error(w, "Failed to save book", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// GET /checkout/{id} (Payment Page)
	mux.HandleFunc("GET /checkout/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		var id int
		fmt.Sscanf(idStr, "%d", &id)

		book, found, err := bookStore.GetByID(id)
		if err != nil || !found {
			http.NotFound(w, r)
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
		// 1. Parse Form
		// In a real app, strict validation would happen here.
		// For now, we trust the HTML5 "required" attributes for the happy path.

		// 2. "Process" Payment (Mock)
		time.Sleep(1 * time.Second) // Simulate network delay

		// 3. Redirect to Home with Success (would be a success page in real app)
		http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
	})

	// Serve Static Assets
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("cmd/web-demo/assets"))))

	// 4. START
	port := ":8082"
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("🚀 Professional Demo running on http://localhost%s\n", port)
	fmt.Printf("--------------------------------------------------\n")
	log.Fatal(http.ListenAndServe(port, mux))
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
		log.Printf("Template Parse Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 3. Execute "base" (because base.html defines the outline)
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("Template Execute Error: %v", err)
	}
}
