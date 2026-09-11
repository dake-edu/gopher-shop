package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// normalizeTitle возвращает однострочное название без краевых пробелов.
// Лимит измеряется в рунах, а не в графемах или байтах.
func normalizeTitle(title string) (string, bool) {
	if !utf8.ValidString(title) {
		return "", false
	}
	title = strings.TrimSpace(title)
	if title == "" || utf8.RuneCountInString(title) > 80 {
		return "", false
	}
	for _, letter := range title {
		if unicode.IsControl(letter) {
			return "", false
		}
	}
	return title, true
}
