package main
import "fmt"
// Go has no pointer arithmetic.

func pointerDemo() {
	// A pointer holds the memory address of a value.
	var p *int
	// The & operator generates a pointer to its operand.
	i := 42
	p = &i
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p)
	// adding & infront of the variable will give us its address
	fmt.Println(&p)
	// set i through the pointer p
	*p = 21
	fmt.Println(*p)
}