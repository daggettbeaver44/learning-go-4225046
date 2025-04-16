package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Printf("Sum is: %v\n\n", sum)

	circleRadius := 15.5
	circlesquare := circleRadius * 2 * math.Pi
	fmt.Printf("Square is: %.2f\n", circlesquare)

}
