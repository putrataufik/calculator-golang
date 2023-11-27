package Mathops

import (
	"errors"
	"math"
)

func Addition(a, b float64) float64 {
	return a + b
}

func Subtraction(a, b float64) float64 {
	return a - b
}

func Multiplication(a, b float64) float64 {
	return a * b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Power(base, exponent float64) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("undefined result: zero to the power of a negative number")
	}

	result := math.Pow(base, exponent)
	if math.IsNaN(result) {
		return 0, errors.New("undefined result: negative base to a non-integer power")
	}

	return result, nil
}
