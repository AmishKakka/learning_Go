// int- stores integers (whole numbers), such as 123 or -123
// float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
// string - stores text, such as "Hello World". String values are surrounded by double quotes
// bool- stores values with two states: true or false

package main
import "fmt"

// var	Can be used inside and outside of functions
// := 	Can only be used inside functions

func variablesDemo() {
	var x int = 10
	fmt.Println(x)

	// simple variable assignments
	var a, b = 3, "Hello"
	fmt.Println(a, b)
	c := 9
	fmt.Println("c: ", c)

	// If a variable should have a fixed value that cannot be changed, use the const keyword.
	const pi = 3.14
	fmt.Println("pi: ", pi)
}

func swap(a string, b string) (string, string) {
	return b, a
}