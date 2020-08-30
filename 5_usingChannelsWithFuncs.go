package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// It is important to mention that foo has to be launched concurrently.
	// main go routine will continue running and get to bar, where the value placed in c will be pulled off and the channel will not block the program.
	go foo(c)

	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// You have to close the channel in order to be able to range over it later.
	// Otherwise the range statement will keep trying to pull from the channel indefinitely.
	close(c)
}

// receive
func bar(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
