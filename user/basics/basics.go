package main

import (
	"fmt"
	"math"
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


// Return multiple values from a single function
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

// Function panicTwo takes in the int i and panics if i is greater than 3, otherwise it calls itself with an incremented argument. Panic stops the ordinary flow of control, whereas recover resumes control of a panicking goroutine.
func panicOne() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling  panicTwo...")
	panicTwo(0)
	fmt.Println("Returned normally from panicTwo")
}

func panicTwo(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in panicTwo", i)
	fmt.Println("Printing panicTwo", i)
	panicTwo(i + 1)
}

// Passing in the star means that we're passing in the pointer to the place in memory, and not the value stored there.
func changeXVal(x *int) {
	*x = 2
}

//Define a struct with the following attributes
type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

//Add a method called area onto the struct which returns a float
func (rect1 *Rectangle) area() float64 {
	return rect1.width * rect1.height
}

//Define an interface. Interfaces are named collections of method signatures

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

// Function main serves as entry point for binary executable
func main() {
	fmt.Println(Sqrt(784)) //28
	fmt.Println(returnMultipleValues(5)) //6, 7, 8
	fmt.Println(factorial(5)) // 120
	fmt.Println(safeDivision(5, 0))
	fmt.Println(safeDivision(6, 3))
	panicOne()
	x := 0
	fmt.Println("x = ", x)
	changeXVal(&x)
	fmt.Println("x = ", x)
	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 20}
	rect2 := Rectangle{0, 50, 10, 20}
	fmt.Println("Rectangle is ", rect1.width, "wide")
	fmt.Println("Rectangle2 is ", rect2.height, "tall")
	fmt.Println("Area of Rectangle is: ", rect1.area())
	circ := Circle{5}
	fmt.Println("Cirlce area = ", getArea(circ))
}


