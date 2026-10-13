package main

import "fmt"

func sum(nunms ...int) int {
	total := 0
	for _, num := range nunms {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
}
