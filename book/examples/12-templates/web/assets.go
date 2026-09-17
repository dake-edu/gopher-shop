// Package assets contains the public presentation files embedded in the binary.
package assets

import "embed"

//go:embed templates static
var Files embed.FS

//go:embed static/css/style.css
var Style string

//go:embed static/brand/gopher-cart.png
var Logo []byte

//go:embed static/brand/favicon.ico
var Favicon []byte

//go:embed static/brand/apple-touch-icon.png
var TouchIcon []byte
