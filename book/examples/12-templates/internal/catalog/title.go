package catalog

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// NormalizeTitle возвращает однострочное название без краевых пробелов.
// Лимит измеряется в рунах, а не в графемах или байтах.
func NormalizeTitle(title string) (string, bool) {
	if !utf8.ValidString(title) {
		return "", false
	}
	title = strings.TrimSpace(title)
	if title == "" || utf8.RuneCountInString(title) > 80 {
		return "", false
	}
	for _, letter := range title {
		if unicode.IsControl(letter) || letter == '\u2028' || letter == '\u2029' {
			return "", false
		}
	}
	return title, true
}
