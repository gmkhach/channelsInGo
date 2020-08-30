package main

import (
    "fmt"
)

func main() {
	c := make(chan int)
	
	// This code will not run because channels block!
	// The following line blocks the current goroutine and grows a fatal error indicating that no go routines are running
	c <- 42
	
	fmt.Println(<- c)
}