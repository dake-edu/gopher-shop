package models

import "testing"

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
