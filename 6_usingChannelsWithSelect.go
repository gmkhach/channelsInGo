package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v, ok := <-q:
			if ok {
				fmt.Println("From the quite channel:", v)
				return
			}
		}
	}

	fmt.Println(<-q)
}