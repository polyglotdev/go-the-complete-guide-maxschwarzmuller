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
	go say("fuck", c)
	for i := range c {
		fmt.Println(i)
	}
}
