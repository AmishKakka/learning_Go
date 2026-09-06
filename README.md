# learning_Go

Each `.go` file focuses on one topic, while all files belong to the same `main` package and are compiled together when the project is run as a package.

## What this covers?

- Variables, constants, type inference, and multiple return values
- Functions, typed parameters, anonymous functions, and closures
- Arrays, slices, `make`, `append`, and two-dimensional collections
- `for` loops, `range`, while-style loops, and recursion
- Structs, maps, and anonymous struct slices
- Pointers and address/value operators
- Methods with value and pointer receivers
- Interfaces and implicit interface implementation
- The empty interface (`interface{}`), also known as `any`

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

The default output demonstrates:

```text
Hello, i am Macintosh
(<nil>, <nil>)
(42, int)
(Hello, string)
```

The calls are intentionally kept in `main.go` as a simple switchboard for exploring the examples one at a time.

## File Guide

| File | Topics and examples |
| --- | --- |
| [main.go](main.go) | Program entry point, package-wide execution, interface values, and type printing with `%T`. |
| [variables.go](variables.go) | Basic types, `var`, short declaration (`:=`), constants, and swapping values with multiple returns. |
| [functions.go](functions.go) | Function declarations, typed parameters, no function overloading, closures, stateful counters, an adder, and Fibonacci generation. |
| [arrays.go](arrays.go) | Fixed-length arrays, inferred array lengths (`...`), indexing, slices, `append`, and a two-dimensional `uint8` collection built with `make`. |
| [loops.go](loops.go) | Go's only looping keyword (`for`), counter loops, `range`, while-style loops, and recursive factorial calculation. |
| [structures.go](structures.go) | Struct fields, maps, deletion and updates, map lookups with the `value, ok` pattern, iteration, and anonymous structs. |
| [pointers.go](pointers.go) | Pointer declarations, addresses with `&`, dereferencing with `*`, and updating a value through a pointer. Go does not support pointer arithmetic. |
| [methods.go](methods.go) | Methods on structs, value receivers, pointer receivers, and a small `Vertex` example. |
| [interfaces.go](interfaces.go) | `Rect` and `Circle` types, the `Shape` interface, and implicit implementation through an `Area() float32` method. |
