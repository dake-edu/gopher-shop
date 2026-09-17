package web

import (
	"html/template"
	"time"

	assets "github.com/dake-edu/gopher-shop/book/examples/12-templates/web"
)

// Each page gets its own set: identical content names cannot replace another page.
// Call only during handler construction, before serving concurrent requests.
func pageTemplate(name string) *template.Template {
	active := "account"
	if name == "home" || name == "workshop" {
		active = "catalog"
	}
	if name == "cart" || name == "order" || name == "payment" {
		active = "cart"
	}
	if name == "cart" {
		name = "home"
	}
	return template.Must(template.New("base").Funcs(template.FuncMap{
		"nav_active": func(section string) bool { return section == active },
		"money":      money,
		"year":       func() int { return time.Now().Year() },
		"brand_name": func() string { return "eGopher" },
	}).ParseFS(
		assets.Files, "templates/layouts/base.html", "templates/partials/*.html",
		"templates/pages/"+name+".html",
	))
}

var styleSource = assets.Style
