package main

import "testing"

func TestSliceSharesArray(t *testing.T) {
	original := []string{"Go", "SQL"}
	part := original[:1]
	part[0] = "HTTP"
	if original[0] != "HTTP" {
		t.Fatal("изменение общего элемента не видно")
	}
}

func TestCopyCreatesIndependentStringElements(t *testing.T) {
	original := []string{"Go", "SQL"}
	independent := make([]string, len(original))
	if copied := copy(independent, original); copied != 2 {
		t.Fatalf("скопировано %d элементов, хотели 2", copied)
	}
	independent[0] = "HTTP"
	if original[0] != "Go" {
		t.Fatal("замена строки копии изменила оригинал")
	}
}

func TestArrayValueCopy(t *testing.T) {
	original := [3]string{"PDF", "EPUB", "HTML"}
	independent := original
	independent[0] = "TXT"
	if original[0] != "PDF" {
		t.Fatal("замена строки в копии массива изменила оригинал")
	}
}
