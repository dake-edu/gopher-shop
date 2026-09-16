package main

import (
	"errors"
	"testing"
)

func TestDeferOrder(t *testing.T) {
	order := 0
	f := func() {
		defer func() { order = order*10 + 1 }()
		defer func() { order = order*10 + 2 }()
	}
	f()
	if order != 21 {
		t.Fatalf("получили %d, хотели 21", order)
	}
}

func TestDeferArgumentAndClosure(t *testing.T) {
	argument := 0
	closure := 0
	f := func() {
		n := 1
		defer func(value int) { argument = value }(n)
		defer func() { closure = n }()
		n = 2
	}
	f()
	if argument != 1 || closure != 2 {
		t.Fatalf("получили %d и %d, хотели 1 и 2", argument, closure)
	}
}

func TestDeferNamedResult(t *testing.T) {
	closeErr := errors.New("закрытие не удалось")
	f := func() (err error) {
		defer func() { err = closeErr }()
		return nil
	}
	if !errors.Is(f(), closeErr) {
		t.Fatal("отложенное действие не изменило результат")
	}
}
