package main

import "fmt"

func main() {
	var result string
	if theAnswer := -42; theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}
