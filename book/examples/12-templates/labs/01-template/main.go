package main

import (
	"fmt"
	"html/template"
	"os"
)

type Page struct{ Title string }

func run() error {
	page, err := template.New("card").Parse("<h1>{{.Title}}</h1>\n")
	if err != nil {
		return err
	}
	return page.Execute(os.Stdout, Page{Title: "<Go>"})
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
