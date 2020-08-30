package main

import (
    "fmt"
)

func main() {
	c := make(chan int)
	
	// Here we create a concurrent goroutine that will be block once 42 is passed to c, but the main goroutine will continue to run.
	go func(){
		c <- 42
	}()
	
	// When the main goroutine reachest this line the value 42 is taken off of channel c and the goroutine is unblocked
	fmt.Println(<- c)
}