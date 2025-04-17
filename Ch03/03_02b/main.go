package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}

	value1 := 42.13
	pointer := &value1
	*pointer = *pointer / 13
	fmt.Println("value1 is:", *pointer)
	fmt.Println("Original value:", value1)
}
