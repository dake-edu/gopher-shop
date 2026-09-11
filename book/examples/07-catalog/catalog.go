package main

import "strings"

// catalogTitles сохраняет порядок ids и отказывает при неполных данных.
func catalogTitles(ids []string, titles map[string]string) ([]string, bool) {
	ordered := make([]string, 0, len(ids))
	for _, id := range ids {
		title, ok := titles[id]
		if !ok || strings.TrimSpace(title) == "" {
			return nil, false
		}
		ordered = append(ordered, title)
	}
	return ordered, true
}
