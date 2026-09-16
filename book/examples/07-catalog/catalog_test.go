package main

import (
	"strings"
	"testing"
)

func TestCatalogOrder(t *testing.T) {
	got, ok := catalogTitles([]string{"b", "a"}, map[string]string{"a": "А", "b": "Б"})
	if !ok || len(got) != 2 {
		t.Fatalf("неожиданный результат: %v, %t", got, ok)
	}
	if got[0] != "Б" || got[1] != "А" {
		t.Fatalf("порядок изменился: %v", got)
	}
}

func TestMissingTitle(t *testing.T) {
	got, ok := catalogTitles([]string{"a", "missing"}, map[string]string{"a": "А"})
	if ok || got != nil {
		t.Fatalf("неполный каталог принят: %v, %t", got, ok)
	}
}

func TestBlankCatalogTitle(t *testing.T) {
	_, ok := catalogTitles([]string{"a"}, map[string]string{"a": "  "})
	if ok {
		t.Fatal("название из пробелов принято")
	}
}

func TestEmptyCatalog(t *testing.T) {
	got, ok := catalogTitles(nil, nil)
	if !ok || len(got) != 0 {
		t.Fatalf("пустой каталог: %v, %t", got, ok)
	}
}

func TestNilMapRead(t *testing.T) {
	_, ok := catalogTitles([]string{"a"}, nil)
	if ok {
		t.Fatal("несуществующее название найдено в nil map")
	}
}

func TestCatalogRetainsTitleRules(t *testing.T) {
	got, ok := catalogTitles([]string{"a"}, map[string]string{"a": "  Go  "})
	if !ok || len(got) != 1 {
		t.Fatalf("допустимый каталог отклонён: %v, %t", got, ok)
	}
	if got[0] != "Go" {
		t.Fatalf("название не подготовлено: %q", got[0])
	}
	for _, title := range []string{"Go\nShop", "Go\u2028Shop", "Go\xff", strings.Repeat("Ә", 81)} {
		got, ok := catalogTitles([]string{"a", "b"}, map[string]string{"a": "Go", "b": title})
		if ok || got != nil {
			t.Fatalf("неверное название %q принято или выдан частичный каталог: %v", title, got)
		}
	}
}
