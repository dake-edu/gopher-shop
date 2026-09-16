package main

import (
	"strings"
	"testing"
)

func TestNormalizeTitle(t *testing.T) {
	got, ok := normalizeTitle("  Шаңырақ Go  ")
	if !ok || got != "Шаңырақ Go" {
		t.Fatalf("получили (%q, %t), хотели (%q, true)", got, ok, "Шаңырақ Go")
	}
}

func TestBlankTitle(t *testing.T) {
	_, ok := normalizeTitle(" \t\n ")
	if ok {
		t.Fatal("пустое после удаления пробелов название принято")
	}
}

func TestTitleRuneBoundary(t *testing.T) {
	_, ok := normalizeTitle(strings.Repeat("Ә", 80))
	if !ok {
		t.Fatal("80 рун должны помещаться, хотя байтов больше")
	}
	_, ok = normalizeTitle(strings.Repeat("Ә", 81))
	if ok {
		t.Fatal("81 руна превышает предел")
	}
}

func TestInvalidUTF8Title(t *testing.T) {
	_, ok := normalizeTitle("Go\xff")
	if ok {
		t.Fatal("неверный UTF-8 принят")
	}
}

func TestMultilineTitle(t *testing.T) {
	_, ok := normalizeTitle("Go\nмагазин")
	if ok {
		t.Fatal("внутренний перевод строки принят")
	}
}

func TestUnicodeLineSeparators(t *testing.T) {
	_, ok := normalizeTitle("Go\u2028магазин")
	if ok {
		t.Fatal("внутренний разделитель строки принят")
	}
	_, ok = normalizeTitle("Go\u2029магазин")
	if ok {
		t.Fatal("внутренний разделитель абзаца принят")
	}
	got, ok := normalizeTitle("\u2028Go\u2029")
	if !ok || got != "Go" {
		t.Fatalf("краевые пробелы должны удаляться: %q, %t", got, ok)
	}
}
