package Mathops

import (
	"errors"
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(5, 3)
	expected := 8.0

	if result != expected {
		t.Errorf("Addition result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 3)
	expected := 2.0

	if result != expected {
		t.Errorf("Subtraction result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(5, 3)
	expected := 15.0

	if result != expected {
		t.Errorf("Multiplication result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(10, 2)
	expected := 5.0

	if result != expected || err != nil {
		t.Errorf("Division result was incorrect, got: %f, want: %f, error: %v", result, expected, err)
	}
	_, zeroErr := Division(5, 0)
	expectedErr := errors.New("division by zero not allowed")

	if zeroErr == nil || zeroErr.Error() != expectedErr.Error() {
		t.Errorf("Division by zero error was not caught correctly, got: %v, want: %v", zeroErr, expectedErr)
	}
}
