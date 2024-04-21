# Maps in Go

Maps are used to group related data together. A map is an unordered collection of key-value pairs, where each key is unique. Maps are used to look up a value by its associated key.

```go
package main

import "fmt"

func main() {
  // Create a map
  m := make(map[string]int)

  // Add key-value pairs to the map
  m["one"] = 1
  m["two"] = 2
  m["three"] = 3

  // Access a value by its key
  fmt.Println(m["two"]) // Output: 2

  // Update a value
  m["two"] = 22
  fmt.Println(m["two"]) // Output: 22

  // Delete a key-value pair
  delete(m, "two")
  fmt.Println(m["two"]) // Output: 0

  // Check if a key exists
  value, ok := m["two"]
  fmt.Println(value, ok) // Output: 0 false

  // Iterate over the map
  for key, value := range m {
      fmt.Println(key, value)
  }
}
```

## `make` function

The `make` built-in function allocates and initializes an object of type **slice, map, or chan (only)**. Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type.

The `make` function is used to create a map. The `make` function takes the type of the map as the first argument and returns a new map.

```go
m := make(map[string]int)
```

## Add key-value pairs

To add a key-value pair to a map, use the following syntax:

```go
m["key"] = value
```

## Access a value

To access a value by its key, use the following syntax:

```go
value := m["key"]
```
