package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Go shop")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	request := httptest.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)
	fmt.Println(response.Code)
	fmt.Print(response.Body.String())
	missing := httptest.NewRecorder()
	mux.ServeHTTP(missing, httptest.NewRequest("GET", "/missing", nil))
	fmt.Println(missing.Code)
}
