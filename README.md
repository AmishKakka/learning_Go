# learning_Go

Each `.go` file focuses on one topic, while all files belong to the same `main` package and are compiled together when the project is run as a package.

## What this covers?

- Variables, constants, type inference, and multiple return values
- Functions, anonymous functions, closures, and generic functions
- Arrays, slices, `make`, `append`, and two-dimensional collections
- `for` loops, `range`, while-style loops, and recursion
- Structs, maps, and anonymous struct slices
- Pointers and address/value operators
- Methods with value and pointer receivers
- Interfaces, implicit implementation, and type assertions
- Empty interfaces (`interface{}`), also called `any`
- The `error` interface and custom error types
- `io.Reader`, custom readers, and the decorator pattern
- Goroutines, `sync.WaitGroup`, channels, buffering, and channel closing
- Generic linked-list types

## Requirements

- Go 1.27.1 or a compatible newer Go release

Check your installed version with:

```bash
go version
```

## Run the Examples

From the repository root, run the complete package:

```bash
go run .
```

Because the examples share the `main` package, `go run .` is preferred over running only `main.go`. Running a single file can omit functions defined in the other files.

## File Guide

| File | Topics and examples |
| --- | --- |
| [main.go](main.go) | Program entry point, package-wide execution, `io.Reader` usage, ROT13 reader composition, type assertions, goroutines, `WaitGroup`, channels, and type printing with `%T`. |
| [variables.go](variables.go) | Basic types, `var`, short declaration (`:=`), constants, and swapping values with multiple returns. |
| [functions.go](functions.go) | Function declarations, typed parameters, no function overloading, closures, stateful counters, an adder, and Fibonacci generation. |
| [arrays.go](arrays.go) | Fixed-length arrays, inferred array lengths (`...`), indexing, slices, `append`, and a two-dimensional `uint8` collection built with `make`. |
| [loops.go](loops.go) | Go's only looping keyword (`for`), counter loops, `range`, while-style loops, and recursive factorial calculation. |
| [structures.go](structures.go) | Struct fields, maps, deletion and updates, map lookups with the `value, ok` pattern, iteration, and anonymous structs. |
| [pointers.go](pointers.go) | Pointer declarations, addresses with `&`, dereferencing with `*`, and updating a value through a pointer. Go does not support pointer arithmetic. |
| [methods.go](methods.go) | Methods on structs, value receivers, pointer receivers, and a `Vertex` example. |
| [interfaces.go](interfaces.go) | `Rect` and `Circle` types, the `Shape` interface, and implicit implementation through an `Area() float32` method. |
| [errors.go](errors.go) | The built-in `error` interface, Newton's-method square root calculation, `ErrNegativeSqrt`, and a time/location-aware `MyError`. |
| [index.go](index.go) | The generic `where[T comparable]` lookup function and a generic `LinkedList[T any]` node type. |
| [goroutines.go](goroutines.go) | Goroutines, scheduling, `WaitGroup`, deferred cleanup with `defer`, channel-based sums, buffered channels, producers, and `close`. |
