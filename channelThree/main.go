package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	go foo(ch)
	bar(ch)
	fmt.Printf("Putting value in to the channel -> %d\n", 100)
	ch <- 100
	fmt.Printf("Pulling value out of the channel -> %d\n", <-ch)
	fmt.Printf("Exiting main routine....\n")
}

func foo(ch chan<- int) {
	fmt.Printf("Sending value into the channel -> %d\n", 10)
	ch <- 10
}

func bar(ch <-chan int) {
	fmt.Printf("Gooooogle me \n")
	fmt.Printf("Recieved Value -> %d\n", <-ch)
}
