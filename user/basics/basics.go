package main

import (
	"fmt"
)

// Iimplement a while loop to calculate the square root of an input number, using Newton's method.
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


// Return two values from a single function
func returnMultipleValues(num int) (int, int, int) {
	return num+1, num+2, num+3
}

// Implement recursive function
func factorial(num int) int {
	if num <= 0 {
		return 1
	} else {
		return num * factorial(num - 1)
	}
}

// The defer keyword delays invocation of a function until all of the surrounding functions within its lexical scope have been invoked. Recover allows execution to continue even if a fatal error appears. Here, that error message is being logged, but it isn't necessary. NOTE: expresion in defer must be funcion call, hence the IIFE-like syntax.
func safeDivision(x, y int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := x / y
	return solution
}




// Function main serves as entry point for binary executable
func main() {
	fmt.Println(Sqrt(784)) //28
	fmt.Println(returnMultipleValues(5)) //6, 7, 8
	fmt.Println(factorial(5)) // 120
	fmt.Println(safeDivision(5, 0))
	fmt.Println(safeDivision(6, 3))
}


