package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 2, 6, 8, 9}
	fmt.Println(numbers)

}
