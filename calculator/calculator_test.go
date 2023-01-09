package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiplySimple(t *testing.T) {
	t.Parallel()
	var want float64 = 10
	got := calculator.Multiply(5, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TestMultiplyNegativeNumberPair(t *testing.T) {
	t.Parallel()
	var want float64 = 10
	got := calculator.Multiply(-5, -2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiplyByZero(t *testing.T) {
	t.Parallel()
	var want float64 = 0
	got := calculator.Multiply(10, 0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideSimple(t *testing.T) {
	t.Parallel()
	var want float64 = 5
	got, _ := calculator.Divide(10, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TestDivideByZero(t *testing.T) {
	t.Parallel()
	var want float64 = 0
	got, _ := calculator.Divide(10, 0)
	if got != 0 {
		t.Errorf("got %f, expected 0", want)
	}
}
