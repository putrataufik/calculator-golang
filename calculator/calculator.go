package calculator

import (
	"calculator-golang/Mathops"
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
	return Mathops.Power(base, exponent)
}
