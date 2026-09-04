package main
import "fmt"

func loopsDemo() {
	// similar sytanx like C++
	for i:=0; i < 10; i+=2 {
		fmt.Println(i)
	}

	fruits := [...]string{"banana", "apple", "guava"}
	// or we can also use: range, similar to Python
	for idx, val := range fruits { 
		fmt.Println(idx, val)
	}

	count := 1
	// this acts as a 'while' loop
	for count <= 5 {
		fmt.Println(count)
		count += 1
	}
}

func factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return x * factorial(x - 1)
}