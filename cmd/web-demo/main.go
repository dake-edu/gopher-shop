package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/store"
)

// PageData creates a standard payload for our templates
type PageData struct {
	Title string
	Books []models.Book
}

func main() {
	// 1. SETUP: Initialize Storage
	//    We use the In-Memory store we created in `internal/store/inmem.go`
	bookStore := store.NewInMemoryBookStore()

	// 2. TEMPLATES: Parse all HTML files
	//    Join parts (header + sidebar + footer + index)
	tmplPath := filepath.Join("cmd", "web-demo", "templates", "*.html")
	if abs, err := filepath.Abs(tmplPath); err == nil {
		fmt.Println("Loading templates from:", abs)
	}

	tmpl, err := template.ParseGlob(tmplPath)
	if err != nil {
		log.Fatalf("Critical: Could not load templates: %v", err)
	}

	// 3. HANDLERS

	// GET / (Home)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		books, err := bookStore.GetAll()
		if err != nil {
			http.Error(w, "Could not fetch books", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title: "Home",
			Books: books,
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Printf("Template Render Error: %v", err)
		}
	})

	// POST /add (Form Submission)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Simple Form Parsing
		var price float64
		fmt.Sscanf(r.FormValue("price"), "%f", &price)

		newBook := &models.Book{
			Title:  r.FormValue("title"),
			Author: r.FormValue("author"),
			Price:  price,
		}

		// ⚓ VISUAL ANCHOR: The Quality Gate
		// Check data before sending to storage
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

	// Serve Static Assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("cmd/web-demo/assets"))))

	// 4. START
	port := ":8082"
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("🚀 Professional Demo running on http://localhost%s\n", port)
	fmt.Printf("--------------------------------------------------\n")
	log.Fatal(http.ListenAndServe(port, nil))
}
