package main

import "fmt"

func main() {
	formats := [3]string{"PDF", "EPUB", "HTML"}
	fmt.Println("Форматы:", formats)

	ids := []string{"go-shop"}
	ids = append(ids, "sql-notes")
	titles := map[string]string{
		"go-shop":   "Go: от первой строки до книжного магазина",
		"sql-notes": "Заметки о SQL",
	}
	ordered, ok := catalogTitles(ids, titles)
	if !ok {
		fmt.Println("Каталог неполон")
		return
	}
	for index, title := range ordered {
		fmt.Println(index+1, title)
	}
	_, found := titles["missing"]
	fmt.Println("Неизвестная книга найдена:", found)
}
