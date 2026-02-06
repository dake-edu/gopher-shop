package main

import (
	"html/template"
	"log"
	"net/http"
)

// Lesson 4: Interaction
// Goal: Accept user input via Forms and Validate it (The Quality Gate).

type Book struct {
	Title  string
	Author string
}

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
	tmpl, _ := template.New("index").Parse(htmlTmpl)

	// GET: Show Page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, books)
	})

	// POST: Handle Data
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")

		// 🛡️ THE QUALITY GATE (Validation)
		// Why? Never trust user input. It could be empty, malicious, or wrong.
		// We reject bad data BEFORE it touches our database.
		if title == "" {
			http.Error(w, "Title cannot be empty!", http.StatusBadRequest)
			return
		}

		// If good, add to shelf
		books = append(books, Book{title, author})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("🚀 Shop v4 running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
