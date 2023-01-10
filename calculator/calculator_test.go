package calculator_test

import (
	"calculator"
	"math"
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
	// So any test, and any test case within each test can be run at
	// the same time, speeds up running the overall collection of tests
	// within the test file hugely. Unfortunately within IntelliJ, the
	//GO plugin doesn't count test cases, just test methods, so you loose
	// visibility of the overall number of tests if you use test cases
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 0, b: 5, want: 5},
		{a: 0, b: 0, want: 0},
	}

	// For each item in slice, get both the index and element, see
	// https://gobyexample.com/range, and ignore the index variable
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
		{a: 1, b: 0, want: math.Inf(1)},
		{a: -1, b: 0, want: math.Inf(-1)},
	}

	for _, tc := range testCases {
		got := calculator.Divide(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
