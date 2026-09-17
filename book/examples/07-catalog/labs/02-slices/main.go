package main

import "fmt"

func main() {
	titles := []string{"Go", "SQL"}
	same := titles
	same[0] = "Web"
	fmt.Println(titles[0])
	independent := make([]string, len(titles))
	copied := copy(independent, titles)
	independent[0] = "Other"
	fmt.Println(copied, titles[0], independent[0])
	empty := make([]string, 0, 3)
	fmt.Println(len(empty), cap(empty), copy(empty, titles))
	empty = append(empty, "Go")
	fmt.Println(len(empty), empty[0])
	base := [3]string{"Go", "SQL", "Web"}
	part := base[:1]
	grown := append(part, "New")
	fmt.Println(base[1], len(part), len(grown))
}
