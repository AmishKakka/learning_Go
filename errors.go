package main
import (
	"fmt"
	"time"
)

// The error type is a built-in interface.
// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
/*
type error interface {
    Error() string
	}
*/

type ErrNegativeSqrt float64

// attaching Error() method to a custom type
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

type MyError struct {
	when time.Time
	where string
}

// attaching my custom Error() method to my custom error type
func (e MyError) Error() string {
	return fmt.Sprintf("at %v %s", e.when, e.where)
}

func run() error {
	return MyError{
		time.Now(),
		"Arizona",
	}
}