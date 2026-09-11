package main

import "fmt"

type Store struct{ values map[string]string }

func (s Store) Set(key, value string) { s.values[key] = value }
func (s Store) Replace()              { s.values = map[string]string{"new": "value"} }

func main() {
	s := Store{values: make(map[string]string)}
	s.Set("book", "Go")
	fmt.Println(s.values["book"])
	s.Replace()
	_, ok := s.values["new"]
	fmt.Println("Замена поля сохранилась:", ok)
}
