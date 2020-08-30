package main

import (
    "fmt"
)

func main() {
	// This is an example of a buffered channel
	// It doesn't have to have a receiver to take values, and therefore does not block the code
	// You can indicate how many values will the buffered channel take
	c := make(chan int, 2)
	
	c <- 42
	c <- 1984

	// The values are pulled off of buffered channel in first-in-first-out order
	fmt.Println(<- c)
	fmt.Println(<- c)
}