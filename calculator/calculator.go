package calculator

import "errors"

// Add takes two numbers and returns the result of adding
// them together.
func Add(a float64, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a and b
func Multiply(a, b float64) float64 {
	var val float64 = 0
	if a != 0 && b != 0 {
		val = a * b
	}
	return val
}

// Divide takes two numbers a and b, and
// returns the result of dividing a by b
// Divide by zero is not allowed
//
// See https://go.dev/doc/faq#exceptions about why GO
// does not have exceptions
func Divide(a, b float64) (float64, error) {
	var val float64 = 0
	var err error = nil
	if b == 0 {
		err = errors.New("cannot divide by zero")
	} else {
		val = a / b
	}
	return val, err
}
