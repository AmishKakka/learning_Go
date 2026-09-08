package main
import (
	"fmt"
	"time"
	"sync"
)
/* A goroutine is a lightweight thread of execution managed directly by the Go runtime,
rather than by your operating system.
The exact order between "world" and "hello" is non-deterministic (unpredictable).
*/

func say(s string) {
	// just prints string 's' with a sleep duration of 100ms
	for i:=0; i < 5; i++ {
		fmt.Println(s)
		// While one goroutine is asleep, the Go runtime scheduler switches CPU time to the other goroutine.
		time.Sleep(100 * time.Millisecond)
	}
}

// The sync.WaitGroup type in Go provides methods to manage multiple goroutines and ensures that main waits until all goroutines have completed.

/* A WaitGroup is a structure with an internal counter that represents the number of goroutines that are being tracked. 
It has three main methods: Add, Done, and Wait.
*/
func ReliableSay(s string, wg *sync.WaitGroup) {
	// defer: delays the execution of a function until the surrounding function returns
	defer wg.Done()
	for i:=0; i < 5; i++ {
		fmt.Println(s)
		// While one goroutine is asleep, the Go runtime scheduler switches CPU time to the other goroutine.
		time.Sleep(100 * time.Millisecond)
	}
}

/*
Channels are the pipes that connect concurrent goroutines.
They allow one goroutine to send data and another goroutine to receive that data safely, preventing race conditions.
*/
func sum(arr []int, ch chan int) {
	total := 0
	for _, v := range arr {
		total += v
	}
	// channel <- variable; i am sending variable to the defined channel
	ch <- total
}

// The sender can close(ch) to signal that no more values will be sent.
func produce(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	// close() signals the reciever that we're done
	close(ch)
}