package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/catalog"
	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/titlefile"
	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/web"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	check := flag.Bool("check", false, "проверить каталог без запуска сервера")
	flag.Parse()
	title, err := titlefile.Load("title.txt")
	if err != nil {
		return err
	}
	books, ok := catalog.Books(title)
	if !ok {
		return errors.New("каталог не прошёл проверку")
	}
	if *check {
		fmt.Println("Книг в каталоге:", len(books))
		return nil
	}
	server := &http.Server{
		Addr: "127.0.0.1:8080", Handler: web.New(books),
		ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second,
		WriteTimeout: 15 * time.Second, IdleTimeout: time.Minute,
	}
	fmt.Println("Откройте http://127.0.0.1:8080; остановка: Ctrl+C")
	return server.ListenAndServe()
}
