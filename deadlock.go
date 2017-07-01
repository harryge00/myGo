package main

import (

// go fmt puts these in alphbetical order.
)

func main() {
	ch := make(chan int)
	<-ch // 阻塞main goroutine, 信道c被锁
}
