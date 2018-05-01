package main

import "fmt"

func fibonacci(c, quit chan int) { // producer of fib(x)
	x, y := 0, 1
	for {
		select { // choose one that doesn’t block
		case c <- x: // send x on c
			x, y = y, x+y // adjust x for next time around
		case <-quit: // quit will block until something sent
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ { // blocking consumers…
			fmt.Println(<-c) // receive x on c when ready & print
		}
		quit <- 0 // after 10, send on the QUIT chan
	}()
	fibonacci(c, quit) // launch the producer
}
