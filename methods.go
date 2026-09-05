package main
// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.

type Vertex struct {
	x, y float32
}

func distance(v Vertex) float32 {
	return v.x + v.y
}

// By doing this, now whenever i create a new Vertex it will have its own method: Abs()
// Value Receiver: Operates on a copy of the struct. It cannot modify the original struct.
func (v Vertex) Abs() float32 {
	if v.x < 0 {
		v.x = -v.x
	}
	if v.y < 0 {
		v.y = -v.y
	}
	return v.x + v.y
}

// Here, again a new Vertex will have a method: Scale()
// Pointer Receiver: Operates directly on the memory address. It can mutate original values and avoids copying large structs.
func (v *Vertex) Scale(f float32) {
	v.x *= f
	v.y *= f
}