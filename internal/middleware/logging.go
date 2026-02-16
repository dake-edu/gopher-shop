package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter is a wrapper to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// RequestLogger logs the incoming HTTP request and its duration.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the writer to capture status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Process request
		next.ServeHTTP(rw, r)

		// Log results
		duration := time.Since(start)

		// ⚓ PEDAGOGICAL NOTE: Structured Logging
		// Instead of a string, we pass key-value pairs.
		// "msg" is the event. Everything else is context.
		slog.Info("http_request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration", duration.String(),
			"ip", r.RemoteAddr,
		)
	})
}
