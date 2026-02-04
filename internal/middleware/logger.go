package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger records the method, path, status, and duration of each request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture Status Code
		wraptWriter := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		log.Printf("[START] %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(wraptWriter, r)

		duration := time.Since(start)
		log.Printf("[END] %s %s %d %v", r.Method, r.URL.Path, wraptWriter.status, duration)
	})
}

// statusWriter is a wrapper around ResponseWriter to capture the status code.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
