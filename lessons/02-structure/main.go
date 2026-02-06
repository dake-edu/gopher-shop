package main

import (
	"log"
	"net/http"
	"text/template"
)

// Lesson 2: Structure
// Goal: Serve HTML (The Storefront) instead of plain text.

const htmlTmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>The Gopher Shop</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="container py-5">
    <h1>📚 The Gopher Shop</h1>
    <p class="lead">Under Construction...</p>
</body>
</html>
`

func main() {
	// Parse the blueprint (Template)
	tmpl, err := template.New("index").Parse(htmlTmpl)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render the blueprint to the user
		tmpl.Execute(w, nil)
	})

	log.Println("🚀 Shop v2 running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
