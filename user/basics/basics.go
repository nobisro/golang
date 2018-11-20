package main

import (
	"fmt"
)

// Here I implement a while loop to calculate the square root of an input number, using Newton's method.

func Sqrt(x float64) float64 {
	valid := true
	i := 10
	z := float64(1)

	for valid {
		if i == 0 {
			valid = false
		}
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		i -= 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(784))
}


