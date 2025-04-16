package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
	fmt.Println("The valuse is", aNumber)
	stringLength, err := fmt.Println("The value is", aNumber)
	if err == nil {
		fmt.Println("Length is:", stringLength)
	}
	fmt.Printf("value of number: %v\n", aNumber)
	fmt.Printf("type of number: %T\n", aNumber)

}
