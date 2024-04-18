# Go - The Complete Guide

[Course Content](https://www.udemy.com/course/go-the-complete-guide)

## Working with Functions and Values

### Functions

- Functions are the building blocks of a program.
- Functions are used to organize code.
- Functions are used to avoid code duplication.
- Functions are used to make code more readable.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Organizing Code with Packages

- Packages are used to organize code.
- Packages are used to avoid name conflicts.
- Packages are used to make code more readable.

```go
package main
```

All Go programs on the first line of the file must have a package declaration. The package declaration is used to define which package the file belongs to. The package declaration is used to organize code.

```go
package learn
```

The package declaration is used to define which package the file belongs to. The package declaration is used to organize code.

## The Importance of the name `main`

- reserved package keyword `main` is used to create an executable program.
- You can not have more than one `main` package in a Go program.

## Understanding Go modules and Building Go Programs

- Go modules are used to manage dependencies.
- 1 module consists of multiple packages.

## Go Types and Null Values

Go comes with a couple of built-in basic types:

- `bool`: for booleans `true` and `false`
- `int`: Numbers **without** decimal places
- `float64`: Numbers **with** decimal places
- `string`: text values(created with double quotes or backticks)

There are also some more complex types:

- `uint`: unsigned integer(only positive or zero) Unsigned means that the **integer can only represent non-negative numbers**.
  - `uint8`: 8-bit unsigned integer(0 to 255)
  - `uint16`: 16-bit unsigned integer(0 to 65535)
  - `uint32`: 32-bit unsigned integer(0 to 4294967295)
  - `uint64`: 64-bit unsigned integer(0 to 18446744073709551615)
- `int32`: 32-bit signed integer(-2147483648 to 2147483647)
- `int64`: 64-bit signed integer(-9223372036854775808 to 9223372036854775807)
- `rune`: represents a Unicode code point
- `byte`: alias for `uint8`
- `complex64`: complex numbers with float32 real and imaginary parts
- `complex128`: complex numbers with float64 real and imaginary parts
- `uintptr`: an unsigned integer to store the uninterpreted bits of a pointer value
- `nil`: the zero value for pointers, interfaces, maps, slices, channels, and functions

## Null Values

In Go value types come with a "null value" which is the value stored in a variable if no other value has been explicitly assigned to it. The null value for a type is the zero value for that type.

- `0` for numeric types
- `false` for the boolean type
- `""` for strings
- `nil` for pointers, functions, interfaces, slices, channels, and maps
- `0.0` for floating-point numbers
- `0+0i` for complex numbers
- `'\u0000'` for runes
- `[]` for arrays and slices
- `struct{}{}` for structs
- `chan struct{}{}` for channels
