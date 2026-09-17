package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/catalog"
)

type pageData struct {
	Title string
	Books []catalog.Book
}

func money(minor int64) string { return fmt.Sprintf("%d.%02d", minor/100, minor%100) }

func New(books []catalog.Book) http.Handler {
	snapshot := append([]catalog.Book(nil), books...)
	page := pageTemplate("home")
	render := func(w http.ResponseWriter, data pageData) {
		var out bytes.Buffer
		if err := page.Execute(&out, data); err != nil {
			http.Error(w, "Не удалось показать страницу", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = out.WriteTo(w)
	}
	mux := http.NewServeMux()
	registerBrand(mux)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		render(w, pageData{Title: "Каталог", Books: snapshot})
	})
	mux.HandleFunc("GET /assets/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		fmt.Fprint(w, styleSource)
	})
	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		for _, book := range snapshot {
			if book.ID == r.PathValue("id") {
				render(w, pageData{Title: book.Title, Books: []catalog.Book{book}})
				return
			}
		}
		http.NotFound(w, r)
	})
	return mux
}
