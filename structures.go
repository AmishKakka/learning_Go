package main
import "fmt"

// collection of members of different data types, into a single variable
type Person struct {
	name string
	age int
	salary float32
}

func structuresDemo() {
	var p1 Person
	p1.age = 24
	p1.name = "Amish"
	p1.salary = 170000
	fmt.Printf("struct p1: %v\n", p1)
	// structs can also be passed as arguments to functions

	// map is an unordered and changeable collection that does not allow duplicates.
	var car = map[string]string {"brand": "Ford", "model": "Mustang", "year": "2017"}
	fmt.Println(car)

	// remove element
	delete(car, "brand")

	// add/update element
	car["brand"] = "Tesla"
	car["model"] = "S"

	// check if value exists
	val, ok := car["model"]
	fmt.Println(val, ok)

	// iterate over the map
	for k, v := range car {
		fmt.Printf("%v: %v,  ", k, v)
	}
}