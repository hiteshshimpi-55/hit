package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Invalid number of arguments")
		fmt.Println("Usage: calculator <num1> <num2> <operation>")
		return
	}

	a := os.Args[1]
	b := os.Args[2]
	op := os.Args[3]

	num1, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Println("Invalid argument for num1:", a)
		return
	}

	num2, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		fmt.Println("Invalid argument for num2:", b)
		return
	}

	result := calculate(num1, num2, op)
	fmt.Println("Result:", result)
}

func calculate(num1, num2 int64, op string) int64 {
	switch op {
	case "-add":
		return addition(num1, num2)
	case "-sub":
		return subtraction(num1, num2)
	case "-mul":
		return multiplication(num1, num2)
	case "-div":
		return division(num1, num2)
	default:
		fmt.Println("Invalid operation:", op)
		return 0
	}
}

// Basic Math Functions
func addition(num1, num2 int64) int64 {
	return num1 + num2
}

func subtraction(num1, num2 int64) int64 {
	return num1 - num2
}

func multiplication(num1, num2 int64) int64 {
	return num1 * num2
}

func division(num1, num2 int64) int64 {
	if num2 == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return num1 / num2
}
