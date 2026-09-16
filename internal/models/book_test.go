package models

import (
	"math"
	"testing"
)

func TestBook_Validate(t *testing.T) {
	tests := []struct {
		name    string
		book    Book
		wantErr bool
	}{
		{
			name: "Valid Book",
			book: Book{
				Title:  "The Go Programming Language",
				Author: "Alan A. A. Donovan",
				Price:  35.99,
			},
			wantErr: false,
		},
		{
			name: "Missing Title",
			book: Book{
				Title:  "",
				Author: "Author",
				Price:  10.0,
			},
			wantErr: true,
		},
		{
			name: "Missing Author",
			book: Book{
				Title:  "Title",
				Author: "",
				Price:  10.0,
			},
			wantErr: true,
		},
		{
			name: "Negative Price",
			book: Book{
				Title:  "Title",
				Author: "Author",
				Price:  -5.0,
			},
			wantErr: true,
		},
		{
			name: "Zero Price",
			book: Book{
				Title:  "Free Book",
				Author: "Author",
				Price:  0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.book.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRejectNonFinitePriceAndBlankFields(t *testing.T) {
	for _, price := range []float64{math.NaN(), math.Inf(1), math.Inf(-1)} {
		b := Book{Title: "Go", Author: "Author", Price: price}
		if b.Validate() == nil {
			t.Fatalf("accepted non-finite price %v", price)
		}
	}
	for _, b := range []Book{{Title: " ", Author: "Author", Price: 1}, {Title: "Go", Author: "\t", Price: 1}} {
		if b.Validate() == nil {
			t.Fatalf("accepted blank field: %+v", b)
		}
	}
	var b *Book
	if b.Validate() == nil {
		t.Fatal("accepted nil book")
	}
}
