package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// ---------------------------------------------------------
// 1. DATA MODEL (The Blueprint)
// ---------------------------------------------------------
type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

// ---------------------------------------------------------
//  2. IN-MEMORY STORAGE (The Database)
//     We use a global variable for simplicity in this demo.
//     In a real app, this belongs in `internal/store`.
//
// ---------------------------------------------------------
var (
	// The Bookshelf (Slice of Books)
	books = []Book{
		{1, "The Go Gopher", "Rob Pike", 25.00},
		{2, "Concurrency in Go", "Katherine Cox-Buday", 30.00},
	}
	// A Mutex to prevent data races (Multiple people writing at once)
	mu sync.Mutex
	// ID Counter
	nextID = 3
)

// ---------------------------------------------------------
//  3. HTML TEMPLATE (The UI)
//     We embed the HTML directly for a single-file demo.
//
// ---------------------------------------------------------
const htmlTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>The Gopher Shop Demo</title>
    <!-- Bootstrap 5.3 CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <link href="/assets/style.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container py-5">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="text-center mb-4">
                    <img src="/assets/gopher.png" alt="Gopher Shop Logo" height="80">
                    <h1 class="mt-3">The Gopher Shop</h1>
                    <p class="text-muted">A Simple In-Memory Bookstore</p>
                </div>

                <!-- 4. ADD BOOK FORM -->
                <div class="card shadow-sm mb-4">
                    <div class="card-header bg-primary text-white">
                        <h5 class="mb-0">Add a New Book</h5>
                    </div>
                    <div class="card-body">
                        <form action="/add" method="POST" class="row g-3">
                            <div class="col-md-6">
                                <label class="form-label">Title</label>
                                <input type="text" name="title" class="form-control" placeholder="e.g., Intro to Go" required>
                            </div>
                            <div class="col-md-6">
                                <label class="form-label">Author</label>
                                <input type="text" name="author" class="form-control" placeholder="Jane Doe" required>
                            </div>
                            <div class="col-md-4">
                                <label class="form-label">Price ($)</label>
                                <input type="number" step="0.01" name="price" class="form-control" placeholder="19.99" required>
                            </div>
                            <div class="col-12 text-end">
                                <button type="submit" class="btn btn-success">Add Book +</button>
                            </div>
                        </form>
                    </div>
                </div>

                <!-- 5. BOOK LIST (Table) -->
                <div class="card shadow-sm">
                    <div class="card-header">
                        <h5 class="mb-0">Available Books</h5>
                    </div>
                    <div class="card-body p-0">
                        <table class="table table-striped table-hover mb-0">
                            <thead class="table-dark">
                                <tr>
                                    <th>ID</th>
                                    <th>Title</th>
                                    <th>Author</th>
                                    <th>Price</th>
                                </tr>
                            </thead>
                            <tbody>
                                {{range .}}
                                <tr>
                                    <td>#{{.ID}}</td>
                                    <td><strong>{{.Title}}</strong></td>
                                    <td>{{.Author}}</td>
                                    <td>${{.Price}}</td>
                                </tr>
                                {{else}}
                                <tr>
                                    <td colspan="4" class="text-center">No books available yet.</td>
                                </tr>
                                {{end}}
                            </tbody>
                        </table>
                    </div>
                </div>

            </div>
        </div>
    </div>
</body>
</html>
`

// ---------------------------------------------------------
// 4. HANDLERS (The Logic)
// ---------------------------------------------------------

func main() {
	// Parse the template once at startup
	tmpl, err := template.New("index").Parse(htmlTmpl)
	if err != nil {
		log.Fatalf("Template error: %v", err)
	}

	// Route 1: GET / (Show the page)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		// "Render": Combine Template + Data (books) -> HTML
		tmpl.Execute(w, books)
	})

	// Route 2: POST /add (Handle Form Submission)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read Form Data
		title := r.FormValue("title")
		author := r.FormValue("author")
		priceStr := r.FormValue("price")

		// Convert Price (Pedagogical simplification: skipping strict error check)
		var price float64
		fmt.Sscanf(priceStr, "%f", &price)

		// Add to "Database"
		mu.Lock()
		book := Book{
			ID:     nextID,
			Title:  title,
			Author: author,
			Price:  price,
		}
		books = append(books, book)
		nextID++
		mu.Unlock()

		// Redirect back to home
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Serve Static Assets (CSS)
	// Maps "/assets/..." URL to "cmd/web-demo/assets/..." folder
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("cmd/web-demo/assets"))))

	// Start Server
	fmt.Println("🚀 Demo Store running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
