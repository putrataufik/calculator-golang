package ui

import (
	"calculator-golang/calculator"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func StartUI() {
	var (
		num1, num2 float64
		err        error
	)

	// code untuk panggil function input first number
	num1, err = getUserInput("Enter the first number: ")
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	// code untuk panggil function input second number
	num2, err = getUserInput("Enter the second number: ")
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// code untuk panggil funciton input simbol
	operation, err := getOperationInput()
	if err != nil {
		fmt.Println("Invalid operation:", err)
		return
	}

	// panggil function performOperation untuk memasukan parameternya
	result, err := performOperation(num1, num2, operation)
	if err != nil {
		fmt.Println("Error during operation:", err)
		return
	}

	fmt.Printf("Result of %.2f %s %.2f is %.2f\n", num1, operation, num2, result)
}

// ini function unutk input angka pertama dan kedua
func getUserInput(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	return num, nil
}

// ini function untuk input simbol matek (+ , -, /, *)
func getOperationInput() (string, error) {
	var operation string
	fmt.Print("Enter the operation (+, -, *, /): ")
	_, err := fmt.Scanln(&operation)
	if err != nil {
		return "", err
	}

	switch operation {
	case "+", "-", "*", "/":
		return operation, nil
	default:
		return "", errors.New("unsupported operation")
	}
}

// ini function untuk menentukan setiap simbol yang di pilih, akan menjalankan yang mana
func performOperation(num1, num2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return calculator.Add(num1, num2), nil
	case "-":
		return calculator.Subtract(num1, num2), nil
	case "*":
		return calculator.Multiply(num1, num2), nil
	case "/":
		return calculator.Divide(num1, num2)
	default:
		return 0, errors.New("unsupported operation")
	}
}
