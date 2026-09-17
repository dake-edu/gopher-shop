package web

import (
	"net/http"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/web"
)

func registerBrand(mux *http.ServeMux) {
	serveImage(mux, "/assets/brand/gopher-cart.png", "image/png", assets.Logo)
	serveImage(mux, "/favicon.ico", "image/vnd.microsoft.icon", assets.Favicon)
	serveImage(mux, "/assets/brand/apple-touch-icon.png", "image/png", assets.TouchIcon)
}

func serveImage(mux *http.ServeMux, path, contentType string, data []byte) {
	mux.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(data)
	})
}
