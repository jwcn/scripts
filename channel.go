package main

import (
	"fmt"
)

// channel 缓冲区阻塞问题

func main() {
	// 第一种, 会报错 fatal error: all goroutines are asleep - deadlock!
	// 因为 chan 长度只有2,写 3 的时候会阻塞直到死掉
	// 注: 声明 chan 时没给长度 make(chan int), 默认为 无缓冲
	// 解决办法, 增加 chan 的长度或者大于长度的异步写入(见第二种)
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	c2 <- 3
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	// 第二种,
	// c2 := make(chan int, 2)
	// c2 <- 1
	// c2 <- 2

	// go func() {
	// 	c2 <- 3
	// }()
	// fmt.Println(<-c2)
	// fmt.Println(<-c2)
	// fmt.Println(<-c2)

}
