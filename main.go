package main

import (
	"fmt"
	"strings"
)

func sum(nunms ...int) int {
	total := 0
	for _, num := range nunms {
		total += num
	}
	return total
}
func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

func split(name string) (firstName, lastName string) {

	index := strings.Index(name, " ")
	if index == -1 {
		return name, ""
	}
	return name[:index], name[index+1:]
}
func main() {

	val, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	first, last := split("mohamed nabil")
	fmt.Println(first, last)

}
