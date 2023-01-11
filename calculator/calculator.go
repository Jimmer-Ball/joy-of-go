package calculator

import (
	"errors"
	"math"
)

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

// Divide takes two numbers a and b, and optional precision
// and return the result of dividing a by b to the precision
// if provided. Note that divide by zero is not allowed
//
// See https://go.dev/doc/faq#exceptions about why GO
// does not have exceptions
func Divide(a float64, b float64, precision ...uint) (float64, error) {
	var val float64 = 0
	var err error = nil
	if b == 0 {
		err = errors.New("cannot divide by zero")
	} else {
		val = a / b
		if len(precision) == 1 {
			val = roundValue(val, precision[0])
		}
	}
	return val, err
}

func Sqrt(a float64, precision ...uint) (float64, error) {
	var result float64 = 0
	var err error = nil
	if a < 0 {
		err = errors.New("cannot square root a negative number")
	} else {
		result = math.Sqrt(a)
		if len(precision) == 1 {
			result = roundValue(result, precision[0])
		}
	}
	return result, err
}

// See https://gosamples.dev/round-float/ for details
func roundValue(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
