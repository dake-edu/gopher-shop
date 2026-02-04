package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/dake-edu/gopher-shop/internal/config"
	"github.com/dake-edu/gopher-shop/internal/middleware"
	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/store"
)

// ------------------------------------------------------------------------------------------------
// ⚓ VISUAL ANCHOR: THE BIG PICTURE (Layered Arch)
// ------------------------------------------------------------------------------------------------
//
//    [Client]  <-->  [ HTTP Handlers ]  <-->  [ Repository Interface ]  <-->  [ PostgreSQL ]
//
// ------------------------------------------------------------------------------------------------

func main() {
	// 1. CONF: Configure the application
	cfg := config.Load()
	addr := ":" + cfg.Port

	// 2. DB: Connect to PostgreSQL
	//    DSN: Data Source Name
	db, err := sql.Open("postgres", cfg.DB.DSN())
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Printf("Warning: Database ping failed: %v", err)
		log.Println("Continuing... (Is docker-compose up?)")
	} else {
		log.Println("✅ Connected to PostgreSQL!")
	}
	// ⚓ CLEANUP: Ensure DB connection closes on exit
	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()

	// 3. STORE: Choose our storage implementation
	var bookStore store.BookRepository = store.NewPostgresBookStore(db)

	// 4. ROUTER: Define the request handlers
	mux := http.NewServeMux()

	// ⚓ HANDLER: Health Check (Advanced)
	//    Checks if the Database is reachable.
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			log.Printf("Health check failed: %v", err)
			http.Error(w, "Service Unavailable (DB down)", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Inject the store into the handlers using a closure
	mux.HandleFunc("GET /api/books", func(w http.ResponseWriter, r *http.Request) {
		books, err := bookStore.GetAll()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting books: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	mux.HandleFunc("GET /api/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		book, found, err := bookStore.GetByID(id)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting book: %v", err)
			return
		}

		if !found {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "book not found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	})

	// ⚓ HANDLER: Create Book
	//    Input:  POST /api/books JSON(Book)
	//    Output: 201 Created JSON(Book) OR 400 JSON(Error)
	mux.HandleFunc("POST /api/books", func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		// 1. Decode JSON payload
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		// 2. ⚓ VALIDATION: Guard Clauses
		//    Check for errors early and return.
		if err := book.Validate(); err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
			return
		}

		// 3. Create in Store
		if err := bookStore.Create(&book); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error creating book: %v", err)
			return
		}

		// 4. Return Success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	})

	// 5. SERVER: Configure Server
	server := &http.Server{
		Addr: addr,
		// Wrap the mux with our middleware chain
		Handler:      middleware.Recoverer(middleware.Logger(mux)),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 6. LIFECYCLE: Graceful Shutdown
	//    Run server in a separate goroutine so we can listen for signals.
	go func() {
		fmt.Printf("--------------------------------------------------\n")
		fmt.Printf("🐹 Gopher Shop API is running on http://localhost%s\n", addr)
		fmt.Printf("--------------------------------------------------\n")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
