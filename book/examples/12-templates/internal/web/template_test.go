package web

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/catalog"
)

func TestEscapeAndEmptyCatalog(t *testing.T) {
	cases := [][]catalog.Book{
		nil,
		{{ID: "go", Title: "<script>alert(1)</script>", Currency: "KZT"}},
	}
	for _, books := range cases {
		w := httptest.NewRecorder()
		New(books).ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
		body := w.Body.String()
		if w.Code != 200 || strings.Contains(body, "<script>") {
			t.Fatalf("unsafe response: %s", body)
		}
		if len(books) == 0 && !strings.Contains(body, "Каталог пока пуст") {
			t.Fatal("missing empty state")
		}
		if len(books) > 0 && !strings.Contains(body, "&lt;script&gt;") {
			t.Fatal("missing escaped title")
		}
	}
	if money(105) != "1.05" {
		t.Fatal("money format")
	}
}
