package main

import (
    "fmt"
)

func main() {
	// Here is a regular channel
	c := make(chan int, 1)

	// You can indicate that the channel is directional as follows:
	// This is a send-only channel
	soc := make(chan<- int, 1)

	// This one is a receive-only channel
	roc := make(<-chan int,  1)

	c <- 1984
	// Since soc is a send-only channel we can send a value to it
	soc <- 42
	// But we cannot send a value to a receive-only channel	
	// roc <- 42

	// The following line will throw an error because soc is a send-only channel
	// fmt.Println(<- soc)

	// Assining send-only or receive-only channel to a regular channel doesn't work
	// c = soc
	// c = roc
	// However, we can assign a regular chennel to either a send-only or receive-only channel
	roc = c
	fmt.Println(<- roc)	
	
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", soc)
	fmt.Printf("%T\n", roc)
}