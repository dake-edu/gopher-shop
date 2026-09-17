package catalog

import "testing"

func TestPriceAfterDiscount(t *testing.T) {
	got, ok := priceAfterDiscount(249900, 10)
	if !ok || got != 224910 {
		t.Fatalf("получили (%d, %t), хотели (224910, true)", got, ok)
	}
}

func TestNoDiscount(t *testing.T) {
	got, ok := priceAfterDiscount(249900, 0)
	if !ok || got != 249900 {
		t.Fatalf("получили (%d, %t), хотели (249900, true)", got, ok)
	}
}

func TestFullDiscount(t *testing.T) {
	got, ok := priceAfterDiscount(249900, 100)
	if !ok || got != 0 {
		t.Fatalf("получили (%d, %t), хотели (0, true)", got, ok)
	}
}

func TestDiscountRounding(t *testing.T) {
	got, ok := priceAfterDiscount(105, 10)
	if !ok || got != 95 {
		t.Fatalf("получили (%d, %t), хотели (95, true)", got, ok)
	}
}

func TestNegativePrice(t *testing.T) {
	_, ok := priceAfterDiscount(-1, 10)
	if ok {
		t.Fatal("отрицательная цена принята")
	}
}

func TestTooLargePrice(t *testing.T) {
	_, ok := priceAfterDiscount(100000001, 10)
	if ok {
		t.Fatal("цена выше предела принята")
	}
}

func TestNegativeDiscount(t *testing.T) {
	_, ok := priceAfterDiscount(100, -1)
	if ok {
		t.Fatal("отрицательная скидка принята")
	}
}

func TestTooLargeDiscount(t *testing.T) {
	_, ok := priceAfterDiscount(100, 101)
	if ok {
		t.Fatal("скидка выше 100 принята")
	}
}

func TestZeroPrice(t *testing.T) {
	got, ok := priceAfterDiscount(0, 10)
	if !ok || got != 0 {
		t.Fatalf("получили (%d, %t), хотели (0, true)", got, ok)
	}
}

func TestMaximumPrice(t *testing.T) {
	got, ok := priceAfterDiscount(100000000, 100)
	if !ok || got != 0 {
		t.Fatalf("получили (%d, %t), хотели (0, true)", got, ok)
	}
}
