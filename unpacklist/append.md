# Growing the Capacity and Length of a Slice

Go uses a specific formula to increase the capacity of a slice when you append elements to it and it needs to grow.

When the underlying array of a slice is not large enough to accommodate new elements, Go will create a new, larger array and copy the elements of the original slice into it. The capacity of the new array (and therefore the new slice) is determined by a specific growth formula:

- If the current capacity of the slice is less than 1024 elements, then the capacity is doubled.
- If the current capacity of the slice is 1024 elements or more, then the capacity is increased by 25%.

This growth formula is designed to balance memory usage and performance. By doubling the capacity for small slices, Go ensures that append operations are fast for small slices. By increasing the capacity by 25% for large slices, Go helps to prevent excessive memory usage when slices get large.

Here's an example that demonstrates how the capacity of a slice grows when you append elements to it:

```go
package main

import "fmt"

func main() {
  // Start with a slice of length 0 and capacity 0.
  s := make([]int, 0)
  fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

  // Append an element to the slice.
  s = append(s, 1)
  fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

  // Keep appending elements until the capacity reaches 1024.
  for len(s) < 1024 {
    s = append(s, 1)
  }
  fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

  // Append one more element.
  s = append(s, 1)
  fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))
}
```
