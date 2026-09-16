package main

import (
	"github.com/dake-edu/gopher-shop/internal/models"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeBookRequest(t *testing.T) {
	for _, tc := range []struct {
		input string
		valid bool
	}{
		{`{"title":"Go","author":"Author","price":1}`, true},
		{`{"title":"Go"} {"price":2}`, false},
		{`{"unexpected":1}`, false},
		{`{"title":"` + strings.Repeat("a", 64<<10) + `"}`, false},
		{``, false},
	} {
		var book models.Book
		req := httptest.NewRequest("POST", "/api/books", strings.NewReader(tc.input))
		err := decodeBookRequest(httptest.NewRecorder(), req, &book)
		if (err == nil) != tc.valid {
			t.Fatalf("length=%d: error=%v", len(tc.input), err)
		}
	}
}
