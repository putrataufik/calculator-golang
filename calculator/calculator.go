package calculator

import (
	"calculator-golang/Mathops"
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return Mathops.Addition(a, b)
}

func Subtract(a, b float64) float64 {
	return Mathops.Subtraction(a, b)
}

func Multiply(a, b float64) float64 {
	return Mathops.Multiplication(a, b)
}

func Divide(a, b float64) (float64, error) {
	return Mathops.Division(a, b)
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
