package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operations string
	InputA     int
	InputB     int
	MSG        string
}

func (e MathError) Error() string {
	var inputs []string
	if e.Operations == "Division" {
		inputs = append(inputs, fmt.Sprintf("%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("%d", e.InputB))
	}
	if e.Operations == "Multiplication" {
		inputs = append(inputs, fmt.Sprintf("%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("%d", e.InputB))
	}
	if e.Operations == "Subtraction" {
		inputs = append(inputs, fmt.Sprintf("%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("%d", e.InputB))
	}
	if e.Operations == "Addition" {
		inputs = append(inputs, fmt.Sprintf("%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("%d", e.InputB))
	}
	return fmt.Sprintf("Math error in %s (%s): %s", e.Operations, strings.Join(inputs, ","), e.MSG)
}

func sum(nums ...int) (int, error) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total, nil
}
func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, MathError{
			Operations: "Division",
			InputA:     a,
			InputB:     b,
			MSG:        "Cannot divide by zero",
		}
	}
	return a / b, nil
}
func main() {

	val, err := sum(1, 2, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	val2, err2 := safeDivision(1, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(val2)
	}
}
