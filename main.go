package main

import (
	"fmt"
	"mygoapp/mymath"
)

func main() {
	num1 := 10
	num2 := 5

	sum := mymath.Add(num1, num2)
	diff := mymath.Subtract(num1, num2)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
}
