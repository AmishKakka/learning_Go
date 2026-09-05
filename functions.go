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

// Think of 'closure' as a function that has its own private backpack.
// Inside that backpack, the function can pack up and carry around any variables that were sitting nearby when it was born. 
// Even if the surrounding code finishes running and disappears, 
// the function keeps its backpack and can read or modify those variables whenever it wants.

// This function returns another function
func createCounter() func() int {
	// variable inside the backpack
	count := 0
	// return the anonymous function (closure)
	return func() int {
		// This inner function references count, 
		// Go keeps count alive by locking it inside here
		count++
		return count
	}
}

// Here, again we define a closure but this time we add an argument to the inner function
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a + b
		return a
	}
}