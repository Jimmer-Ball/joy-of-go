package calculator_test

import (
	"calculator"
	"testing"
)

/*
 * Use localised test cases rather than copying every test every time
 */
type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	// Allows any test, and any test case within each test to be run at
	// the same time. Speeds up running the overall collection of tests
	// within the test file hugely. Unfortunately within IntelliJ, the
	// GO plugin doesn't count test cases, just test functions, so you
	// loose visibility of the overall number of actual tests run when
	// you use test cases
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 0, b: 5, want: 5},
		{a: 0, b: 0, want: 0},
	}

	// For each item in slice, get both the index and element, see
	// https://gobyexample.com/range, and ignore the index variable
	// using the special blank identifier "_"
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 4, b: 6, want: -2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 5, b: 2, want: 10},
		{a: -5, b: -2, want: 10},
		{a: 10, b: 0, want: 0},
		{a: 0, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 10, b: 2, want: 5},
		{a: -10, b: 2, want: -5},
		{a: 10, b: -2, want: -5},
		{a: 0, b: 5, want: 0},
	}
	for _, tc := range testCases {
		got, _ := calculator.Divide(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideByZero(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 1, b: 0},
		{a: -1, b: 0},
	}
	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			// We expected an error
			t.Log("Expected error returned:", err)
		} else {
			t.Errorf("Divide(%f, %f): Failed to provide expected error", tc.a, tc.b)
		}
	}
}

func TestDivideAndRound(t *testing.T) {
	t.Parallel()

	var a float64 = 1
	var b float64 = 3
	var precision uint = 6
	var want = 0.333333
	got, _ := calculator.Divide(a, b, precision)
	if want != got {
		t.Errorf("DivideAndRound(%f, %f, %d): want %f, got %f", a, b, precision, want, got)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	var a float64 = 9
	var want float64 = 3
	got, _ := calculator.Sqrt(a)
	if want != got {
		t.Errorf("Sqrt(%f): want %f, got %f", a, want, got)
	}
}

func TestSqrtNegative(t *testing.T) {
	t.Parallel()

	var a float64 = -9
	_, err := calculator.Sqrt(a)
	if err != nil {
		// We expected an error
		t.Log(err)
	} else {
		t.Errorf("Sqrt(%f):  Failed to provide expected error", a)
	}
}

func TestSqrtAndRound(t *testing.T) {
	t.Parallel()

	var a float64 = 12
	var want = 3.464102
	var precision uint = 6
	got, _ := calculator.Sqrt(a, precision)
	if want != got {
		t.Errorf("Sqrt(%f): want %f, got %f", a, want, got)
	}
}
