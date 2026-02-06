package main

import (
	"html/template"
	"log"
	"net/http"
)

// Lesson 3: Data (The Brain)
// Goal: Display dynamic data using Structs (Boxes) and Slices (Shelves).

// 1. The Box (Struct)
type Book struct {
	Title  string
	Author string
	Price  float64
}

// 2. The Shelf (Slice)
var books = []Book{
	{"The Go Gopher", "Rob Pike", 25.00},
	{"Concurrency in Go", "Katherine Cox-Buday", 30.00},
}

const htmlTmpl = `
<!DOCTYPE html>
<head>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="container py-5">
    <h1>📚 The Gopher Shop</h1>
    
    <!-- 3. The Conveyor Belt (Range Loop) -->
    <ul class="list-group mt-4">
    {{range .}}
        <li class="list-group-item">
            <strong>{{.Title}}</strong> by {{.Author}} - ${{.Price}}
        </li>
    {{end}}
    </ul>
</body>
</html>
`

func main() {
	tmpl, _ := template.New("index").Parse(htmlTmpl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Pass the "Shelf" (books) to the template
		tmpl.Execute(w, books)
	})

	log.Println("🚀 Shop v3 running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
