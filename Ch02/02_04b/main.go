package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	f1, f2, f3 := 12.1, 54.2, 56.1
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	total := float64(i1) + f2
	fmt.Println("Total is: ", total)

	div := float64(i2) / f3
	fmt.Println("Div: ", div)

	rem := i3 % i1
	fmt.Println("Rem:", rem)
}
