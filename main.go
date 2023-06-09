package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type operationFunc func([]float64) (float64, error)

var operations = map[string]operationFunc{
	"-add":  addition,
	"-sub":  subtraction,
	"-mul":  multiplication,
	"-div":  division,
	"-cm2m": cmToM,
	"-m2cm": mToCm,
	"-l2ml": lToMl,
	"-ml2l": mlToL,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Invalid number of arguments")
		fmt.Println("Usage: calculator <operation> <num1> [<num2>...]")
		return
	}

	op := os.Args[1]
	args := os.Args[2:]

	nums := make([]float64, len(args))
	for i, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("Invalid argument:", arg)
			return
		}
		nums[i] = num
	}

	if operation, ok := operations[op]; ok {
		result, err := operation(nums)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Result:", result)
		}
	} else {
		fmt.Println("Invalid operation:", op)
	}
}

func addition(nums []float64) (float64, error) {
	if len(nums) != 2 {
		return 0, errors.New("Addition requires 2 arguments")
	}
	return nums[0] + nums[1], nil
}

func subtraction(nums []float64) (float64, error) {
	if len(nums) != 2 {
		return 0, errors.New("Subtraction requires 2 arguments")
	}
	return nums[0] - nums[1], nil
}

func multiplication(nums []float64) (float64, error) {
	if len(nums) != 2 {
		return 0, errors.New("Multiplication requires 2 arguments")
	}
	return nums[0] * nums[1], nil
}

func division(nums []float64) (float64, error) {
	if len(nums) != 2 {
		return 0, errors.New("Division requires 2 arguments")
	}
	if nums[1] == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return nums[0] / nums[1], nil
}

func cmToM(nums []float64) (float64, error) {
	if len(nums) != 1 {
		return 0, errors.New("cm to m conversion requires 1 argument")
	}
	return nums[0] / 100, nil
}

func mToCm(nums []float64) (float64, error) {
	if len(nums) != 1 {
		return 0, errors.New("m to cm conversion requires 1 argument")
	}
	return nums[0] * 100, nil
}

func lToMl(nums []float64) (float64, error) {
	if len(nums) != 1 {
		return 0, errors.New("l to ml conversion requires 1 argument")
	}
	return nums[0] * 1000, nil
}

func mlToL(nums []float64) (float64, error) {
	if len(nums) != 1 {
		return 0, errors.New("ml to l conversion requires 1 argument")
	}
	return nums[0] / 1000, nil
}
