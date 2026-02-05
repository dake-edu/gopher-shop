# 14. The Onion (Middleware)

In a professional web server, you need to do things *around* the main logic:
- Log every request.
- Check security headers.
- Handle crashes (Panic Recovery).

## 14.1 Layers of the Onion
**Middleware** wraps your handler like layers of an onion.
The Request must pass through all outer layers to get to the center (Your Logic), and then pass back out.

**Request** $\rightarrow$ [Logger] $\rightarrow$ [Security] $\rightarrow$ **(Handler)** $\rightarrow$ [Security] $\rightarrow$ [Logger] $\rightarrow$ **Response**

## 14.2 The Chainer
In Go, middleware is a function that takes a Handler and returns a Handler.

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w, r) {
        fmt.Println("Request started...")
        next.ServeHTTP(w, r) // Pass to next layer
        fmt.Println("Request finished.")
    })
}
```
