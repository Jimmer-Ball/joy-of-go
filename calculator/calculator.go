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
	var returnValue float64 = 0
	if a == 0 || b == 0 {
		// Avoid rounding errors
		returnValue = 0
	} else {
		returnValue = a * b
	}
	return returnValue
}

// Divide takes two numbers a and b, and
// returns the result of dividing a by b,
// dividing by zero is evil
func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		// Do not divide by zero, as that is evil
		return 0, errors.New("cannot divide by zero, bad piglet")
	}
	return a / b, nil
}
