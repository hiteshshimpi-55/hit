package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := "0"
	b := "0"
	op := "-add"

	if len(os.Args) >= 3 {
		a = os.Args[1]
		b = os.Args[2]
		op = os.Args[3]
	} else {
		fmt.Println("Please pass arguments and operations to perform. . . .")
		return
	}

	num1, _ := strconv.ParseInt(a, 10, 64)
	num2, _ := strconv.ParseInt(b, 10, 64)
	var result int64
	switch op {
	case "-add":
		result = addition(num1, num2)
	case "-sub":
		result = subtraction(num1, num2)
	case "-mul":
		result = multiplication(num1, num2)
	case "-div":
		result = division(num1, num2)
	}

	fmt.Println(result)
}

func addition(num1 int64, num2 int64) int64 {
	result := num1 + num2
	return result
}

func subtraction(num1 int64, num2 int64) int64 {
	result := num1 - num2
	return result
}

func multiplication(num1 int64, num2 int64) int64 {
	result := num1 * num2
	return result
}

func division(num1 int64, num2 int64) int64 {
	result := num1 / num2
	return result
}
