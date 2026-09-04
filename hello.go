// Packages are ways of grouping up related Go code together.
package main

// function name: Hello; then we pass the return type (required) of the function.
func Hello() string {
	return "Hello, i am Macintosh"
}

// Function to add 2 numbers
func add(x int, y int) int {
	// for the arguments in the function, we can also have: x, y int
	return x + y
}

// Go doesn't have function overloading.
func addFloat(x, y float32) float32 {
	return x + y
}