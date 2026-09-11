package main

import "testing"

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
