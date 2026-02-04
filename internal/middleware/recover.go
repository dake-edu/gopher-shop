package middleware

import (
	"log"
	"net/http"
)

// Recoverer catches panics and returns a 500 Internal Server Error.
// This prevents the entire server from crashing due to a single bad request.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
