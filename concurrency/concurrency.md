# Concurrency in Go

Concurrency in Go is a way of structuring software, particularly as a way of writing clean code that interacts well with the real world. It is not parallelism, nor is it about performance, but it enables parallelism and efficient performance when used correctly.

Go provides two main features for concurrent programming:

1. Goroutines: A Goroutine is a lightweight thread managed by the Go runtime. They are very cheap compared to threads, you can create thousands of them at the same time.
2. Channels: Channels are a typed conduit through which you can send and receive values with the channel operator.

Here is an example of how to use goroutines and channels:

```go
package main

import (
	"fmt"
	"time"
)

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- s
	}
	close(c)
}

func main() {
	c := make(chan string)
	go say("world", c)
	for i := range c {
		fmt.Println(i)
	}
}
```

In this example, we create a goroutine with `go say("world", c)`. This goroutine starts executing concurrently with the call. The `say` function sends a string to the channel `c` every 100 milliseconds, and then closes the channel.

In the main function, we use a for loop to receive values from the channel `c`. When the channel is closed and all values have been received, the loop terminates.

This is a simple example of how you can use goroutines and channels to achieve concurrency in Go.
