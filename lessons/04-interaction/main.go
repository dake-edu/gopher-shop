package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Lesson 4: Interaction
// Goal: Accept user input via Forms and Validate it (The Quality Gate).

type Book struct {
	Title  string
	Author string
}

var booksMu sync.Mutex

var books = []Book{
	{"The Go Gopher", "Rob Pike"},
}

const htmlTmpl = `
<!DOCTYPE html>
<head><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet"></head>
<body class="container py-5">
    <h1>📚 The Gopher Shop</h1>

    <!-- FORM -->
    <form action="/add" method="POST" class="mb-4 row g-3">
        <div class="col-auto"><input type="text" name="title" class="form-control" placeholder="Title"></div>
        <div class="col-auto"><input type="text" name="author" class="form-control" placeholder="Author"></div>
        <div class="col-auto"><button type="submit" class="btn btn-primary">Add Book</button></div>
    </form>

    <!-- LIST -->
    <ul class="list-group">
    {{range .}}
        <li class="list-group-item">{{.Title}} ({{.Author}})</li>
    {{end}}
    </ul>
</body>
</html>
`

func main() {
	tmpl := template.Must(template.New("index").Parse(htmlTmpl))

	// GET: Show Page
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		booksMu.Lock()
		snapshot := append([]Book(nil), books...)
		booksMu.Unlock()
		if err := tmpl.Execute(w, snapshot); err != nil {
			log.Print(err)
		}
	})

	// POST: Handle Data
	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		title := strings.TrimSpace(r.PostForm.Get("title"))
		author := strings.TrimSpace(r.PostForm.Get("author"))

		// 🛡️ THE QUALITY GATE (Validation)
		// Why? Never trust user input. It could be empty, malicious, or wrong.
		// We reject bad data BEFORE it touches our database.
		if title == "" || author == "" {
			http.Error(w, "Title and author are required!", http.StatusBadRequest)
			return
		}

		// If good, add to shelf
		booksMu.Lock()
		books = append(books, Book{title, author})
		booksMu.Unlock()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("🚀 Shop v4 running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
