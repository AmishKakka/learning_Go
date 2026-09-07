package main

// By attaching [T comparable] to a function and passing data type which has comaparable function, we can do overloading
// As in have generic functions
func where[T comparable](a []T, target T) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}

// building a linked list with generic type
type LinkedList[T any] struct {
	val T
	next *LinkedList[T]
}