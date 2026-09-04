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

	pointerDemo()
}