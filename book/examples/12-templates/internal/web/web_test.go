package web

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/catalog"
)

func TestRoutes(t *testing.T) {
	books, ok := catalog.Books("Go")
	if !ok {
		t.Fatal("invalid fixture")
	}
	handler := New(books)
	books[0].Title = "Changed outside"
	for _, tc := range []struct {
		method, path string
		code         int
	}{
		{"GET", "/", 200}, {"GET", "/books/go-shop", 200},
		{"GET", "/books/missing", 404}, {"GET", "/missing", 404},
		{"POST", "/books/go-shop", 405},
	} {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		if w.Code != tc.code {
			t.Fatalf("%s %s: %d", tc.method, tc.path, w.Code)
		}
		if strings.Contains(w.Body.String(), "Changed outside") {
			t.Fatal("catalog aliases caller")
		}
	}
}
