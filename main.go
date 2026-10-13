package main

import (
	"fmt"
	"strings"
)

type CustomError struct {
	
}

func split(name string) (firstName, lastName string) {

	index := strings.Index(name, " ")
	if index == -1 {
		return name, ""
	}
	return name[:index], name[index+1:]
}
func main() {


}
