// 斐波那契数列, 练习 channel
package main

import (
	"fmt"
)

func main() {
	out := worker()

	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
}

func worker() chan int {
	x := make(chan int, 2)
	a, b, out := make(chan int, 2), make(chan int, 2), make(chan int, 2)

	go func() {
		for {
			i := <-x

			a <- i
			b <- i
			out <- i
		}
	}()

	go func() {
		x <- 0
		x <- 1
		<-a

		for {
			x <- <-a + <-b
		}
	}()

	return out
}
