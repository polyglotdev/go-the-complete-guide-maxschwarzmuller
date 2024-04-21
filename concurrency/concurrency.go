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
}
