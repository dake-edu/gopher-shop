package web

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/web"
)

func TestBrandAssetsAndClosedDirectory(t *testing.T) {
	mux := http.NewServeMux()
	registerBrand(mux)
	for _, item := range []struct {
		path, kind string
		data       []byte
	}{
		{"/assets/brand/gopher-cart.png", "image/png", assets.Logo},
		{"/favicon.ico", "image/vnd.microsoft.icon", assets.Favicon},
		{"/assets/brand/apple-touch-icon.png", "image/png", assets.TouchIcon},
	} {
		response := httptest.NewRecorder()
		mux.ServeHTTP(response, httptest.NewRequest("GET", item.path, nil))
		if response.Code != 200 || response.Header().Get("Content-Type") != item.kind || !bytes.Equal(response.Body.Bytes(), item.data) {
			t.Fatal("incorrect embedded brand response", item.path, response.Code)
		}
	}
	for _, path := range []string{"/assets/brand/", "/assets/brand/.gitkeep", "/assets/brand/unknown.png"} {
		response := httptest.NewRecorder()
		mux.ServeHTTP(response, httptest.NewRequest("GET", path, nil))
		if response.Code != 404 {
			t.Fatal("unexpected public asset", path, response.Code)
		}
	}
}
