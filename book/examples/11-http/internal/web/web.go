package web

import (
	"fmt"
	"net/http"

	"github.com/dake-edu/gopher-shop/book/examples/11-http/internal/catalog"
)

// New copies the catalog; handlers only read this snapshot.
func New(books []catalog.Book) http.Handler {
	snapshot := make([]catalog.Book, len(books))
	copy(snapshot, books)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		for _, book := range snapshot {
			fmt.Fprintln(w, book.ID, book.Title)
		}
	})
	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		for _, book := range snapshot {
			if book.ID == r.PathValue("id") {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				fmt.Fprintln(w, book.Title)
				return
			}
		}
		http.NotFound(w, r)
	})
	return mux
}
