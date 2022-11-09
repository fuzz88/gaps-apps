package main

import (
	"syscall/js"
)

func fibonacci(c chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
	close(c)
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 42; i++ {
			js.Global().Call("add_number", <-c);
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
