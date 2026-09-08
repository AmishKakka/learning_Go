// Go requires a go.mod file in your project directory to manage its module system
// Command: go mod init folder_name
package main

// the library that allows us to print anything to stdout.
import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	// "time"
)

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
	// r := Rect{3.5, 6.2}
	// c := Circle{3.44}
	// printArea(r)
	// printArea(c)

	// Empty interfaces are used by code that handles values of unknown type.
	// inplace of: interface{}, we can also use keyword: any
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "Hello"
	// describe(i)
	// Type assertion provides access to an interface value's underlying concrete value
	// i.(T); where T can be: string, float32, int 
	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// Custom Errors
	// fmt.Println(sqrt(2))
	// fmt.Println(sqrt(-2))

	// err := run()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Using io.Reader 
	s := strings.NewReader("Hello, World!")

	// Custom Reader
	// r := &MyReader{"Hello, Worldz!", 0}

	// Custom reader that modifies elements by 13
	r := rot13Reader{s}
	// io.Copy reads from r and streams directly to terminal stdout
	io.Copy(os.Stdout, &r)
	fmt.Print("\n")

	// making a byte array of length 8
	// b := make([]byte, 20)
	// for {
	// 	// populate the byte array with content being read
	// 	n, err := r.Read(b)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("n = %v, str = %q, b = %v\n", n, b[:n], b)
	// }

	// Go functions can be written to work on multiple types using type parameters.
	// a1 := []int {10, 2, 3, 45, 98, 1}
	// target := 98
	// fmt.Printf("found value %v at index %v\n", target, where(a1, target))
	// a2 := []string {"five", "foo", "voo", "boo"}
	// str := "hello"
	// fmt.Printf("found value %v at index %v\n", str, where(a2, str))

	// linked list
	// dummy := &LinkedList[int]{0, nil}
	// prev := dummy
	// for i := 1; i <= 5; i++ {
	// 	node := &LinkedList[int]{val:i, next:nil}
	// 	prev.next = node
	// 	prev = node
	// }
	// curr := dummy.next
	// for curr != nil {
	// 	fmt.Printf(" %v -> ", curr.val)
	// 	curr = curr.next
	// }
	// fmt.Print("nil")

	//  --------------------------------------------------------------------------------- 

	// Goroutines
	go say("hello")
	say("world")
	// If i make - say("world") also a goroutine, then nothing will be printed,
	// because main() has no remaining work to do, so it terminates before either background goroutine wakes up from its first sleep.

	// A more reliable way to synchronize goroutines is by using a WaitGroup. 
	var wg sync.WaitGroup
	// Adding new goroutines before they are called
	wg.Add(2)

	go ReliableSay("hello", &wg)
	go ReliableSay("world", &wg)

	// wg.Wait() blocks the main goroutine until the counter reaches 0, ensuring all goroutines have finished before main exits.
	wg.Wait()
	fmt.Println("All go routines completed.")

	// Channels
	// this is an unbuffered channel, meaning no specific length is defined
	ch := make(chan int, 10)
	a := []int {10, -3, 8, 90, 0, 4}
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	// variable <- channel, get value from the channel
	x, y := <-ch, <-ch
	fmt.Printf("x: %v, y: %v, x+y: %v\n", x, y, x + y)
	// Now, channel 'ch' will be empty

	// Buffered channel
	c := make(chan int, 2)
	c <- 100
	c <- 34
	// removing items from the buffer: '<- channel'
	fmt.Println(<- c)
	fmt.Println(<- c)

	// Sending only required amount of data
	// although channel is defined for a length of 10, i am producing only 5 items and reading just them
	go produce(5, ch)
	for val := range ch {
		fmt.Println("recieved: ", val)
	}
}


// Wrapping io.Reader inside another is known as the Decorator pattern
type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	// the underlying reader will loop through the string 
	n, err := rot.r.Read(p)
	// modifying the bytes that were actually read
	for i := 0; i < n; i++ {
		b := p[i]
		if 'A' <= b && b <= 'Z' {
			p[i] = 'A' + (b - 'A' + 13) % 26
		}
		if 'a' <= b && b <= 'z' {
			p[i] = 'a' + (b - 'a' + 13) % 26
		}
	}
	return n, err
}

type MyReader struct {
	s string
	// pos will help us track the current index we are on in the string s
	pos int
}

func (r *MyReader) Read(p []byte) (int, error) {
	// making sure we are within bounds of the given string
	if r.pos >= len(r.s) {
		return 0, io.EOF
	}
	n := 0
	for n < len(p) && r.pos < len(r.s) {
		// replacing letters on even positions
		if r.pos % 2 == 0 {
			p[n] = 'A'
		}else{
			p[n] = r.s[r.pos]
		}	
		r.pos += 1
		n += 1
	}
	return n, nil
}

func describe(i interface{}) {
	// %T in print statement prints out the variable's type
	fmt.Printf("(%v, %T)\n", i, i)
}