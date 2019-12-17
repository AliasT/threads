package main

import (
	"time"
)

func main() {
	ch := make(chan int, 3) // buffered

	for i := 0; i < 3; i++ {
		// golang也有js的作用域变量共享问题
		j := i
		go func() {
			time.Sleep(time.Second)
			ch <- j
		}()
	}

	for {
		select {
		case r := <-ch:
			println(r)

		// 避免死锁报错？
		case <-time.After(2 * time.Second):
			return
		}
	}
}
