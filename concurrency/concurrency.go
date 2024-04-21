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

func greet(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- s
	}
	close(c)
}

// create an example with a channel that is a boolean
func isEven(n int, c chan bool) {
	if n%2 == 0 {
		c <- true
	} else {
		c <- false
	}
	close(c)
}

func main() {
	c := make(chan string)
	go say("hey", c)
	for i := range c {
		fmt.Println(i)
	}

	c1 := make(chan string)
	go greet("there", c1)
	for i := range c1 {
		fmt.Println(i)
	}

	c2 := make(chan bool)
	go isEven(2, c2)
	<-c2
	fmt.Println("2 is even")
}
