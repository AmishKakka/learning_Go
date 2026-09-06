package main

import "fmt"

type Rect struct {
	width, length float32
}

// Value receiver method for Rect struct
func (r Rect) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

// Value receiver method for Circle struct
func (c Circle) Area() float32 {
	return 3.142 * c.radius * c.radius
}

// An interface is a set of method signatures that defines expected behavior.
// Any type that implements an Area() method returning a float32 is a Shape.
// kind of like an Abstract class with its methods in Python
type Shape interface {
	Area() float32
}

func printArea(s Shape) {
	fmt.Printf("Area: %.3f\n", s.Area())
}