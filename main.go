// Go requires a go.mod file in your project directory to manage its module system
// Command: go mod init folder_name
package main
// the library that allows us to print anything to stdout.
import "fmt"

// if i am declaring a main function here, and then again in some other file (hello_world.go) then it will throw - duplicate function declared error.
// Now, to run this we will have to compile all the files in this folder, because we are using function from different files
// So, to compile and run this file: go run .
// otherwise, just ot compile and run a single independent file: go run file_name.go
func main() {
	fmt.Println(Hello())

	// add 2 ints
	// fmt.Println(add(3, 4))

	// add 2 floats
	// fmt.Println(addFloat(9.0, 3.0))

	// Declare a function inline without giving it a name, assign it to a variable, 
	// pass it to other functions, or execute it immediately
	// add := func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(add(4, 5))

	// Here, both A & B will get their own backpack containing 'count' variable
	// This makes closure stateful, it lets you attach persistent state directly to a 
	// function without using global variables or writing a full object-oriented struct.
	// counterA, counterB := createCounter(), createCounter()
	// fmt.Println(counterA())
	// fmt.Println(counterA())
	// fmt.Println(counterB())

	// Again, 'pos' and 'neg' will have their own state of 'sum' variable
	// pos, neg := adder(), adder()
	// for i := range 5 {
	// 	fmt.Println(pos(i), "  ", neg(-2*i))
	// }

	// Fibonacci function
	// fib := fibonacci()
	// fmt.Println("Fibonacci series: ")
	// for range 10 {
	// 	fmt.Println(fib())
	// }

	// defining a variable
	// variablesDemo()

	// var a = "SFO"
	// var b = "PHX"
	// fmt.Println("a: ", a, "b: ", b)
	// fmt.Println(swap(a, b))

	// arraysDemo()

	// loopsDemo()

	// Recursive function
	// fmt.Println(factorial(4))

	// structuresDemo()

	// pointerDemo()

	// Accessing methods attached to a struct
	// v := Vertex{3, -4}
	// fmt.Printf("v: %v\n", v)
	// // printing normal distance
	// fmt.Println(distance(v))
	// // taking only absolute values
	// fmt.Println(v.Abs())
	// // changing original values
	// v.Scale(10)
	// fmt.Printf("v: %v\n", v)

	// Implementing interface
	r := Rect{3.5, 6.2}
	c := Circle{3.44}
	printArea(r)
	printArea(c)

	// Empty interfaces are used by code that handles values of unknown type.
	// inplace of: interface{}, we can also use keyword: any
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

func describe(i interface{}) {
	// %T in print statement prints out the variable's type
	fmt.Printf("(%v, %T)\n", i, i)
}