package main

import (
	"fmt"
	"html/template"
	"os"
)

type Page struct{ Title string }

func run() error {
	source := `{{define "base"}}<main>{{template "content" .}}</main>{{template "footer" .}}
{{end}}{{define "content"}}<h1>{{.Title}}</h1>{{end}}{{define "footer"}}<footer>eGopher</footer>{{end}}`
	page, err := template.New("pages").Parse(source)
	if err != nil {
		return err
	}
	return page.ExecuteTemplate(os.Stdout, "base", Page{Title: "Go"})
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
