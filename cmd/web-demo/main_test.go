package main

import (
	"github.com/dake-edu/gopher-shop/internal/service"
	"github.com/dake-edu/gopher-shop/internal/store"
	"github.com/dake-edu/gopher-shop/internal/worker"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCheckoutUsesFormBookID(t *testing.T) {
	queue := make(chan worker.Order, 1)
	handler := newDemoHandler(service.NewBookService(store.NewInMemoryBookStore()), queue)
	for _, tc := range []struct {
		id     string
		status int
	}{{"bad", 400}, {"1junk", 400}, {"0", 400}, {"999", 404}, {"1", 202}, {"2", 503}} {
		req := httptest.NewRequest(http.MethodPost, "/checkout", strings.NewReader(url.Values{"book_id": {tc.id}}.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		if w.Code != tc.status {
			t.Fatalf("id=%s: got %d, want %d: %s", tc.id, w.Code, tc.status, w.Body.String())
		}
	}
	if len(queue) != 1 {
		t.Fatalf("queued %d jobs", len(queue))
	}
	if job := <-queue; job.ID != 1 {
		t.Fatalf("wrong book: %+v", job)
	}
}
func TestBadPriceDoesNotCreateBook(t *testing.T) {
	repo := store.NewInMemoryBookStore()
	handler := newDemoHandler(service.NewBookService(repo), make(chan worker.Order, 1))
	for _, price := range []string{"1oops", "NaN", "+Inf", "0"} {
		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(url.Values{"title": {"Go"}, "author": {"Author"}, "price": {price}}.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("price=%s: status=%d", price, w.Code)
		}
	}
	books, err := repo.All()
	if err != nil || len(books) != 3 {
		t.Fatalf("invalid form changed catalog: %v %v", books, err)
	}
}
