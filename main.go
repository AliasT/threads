package main

import (
	"sync"
	"time"
)

func main() {
	main1()
	main2()
}

func main1() {
	ch := make(chan int, 3) // buffered
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		// golang也有js的作用域变量共享问题
		j := i
		go func() {
			time.Sleep(time.Second)
			ch <- j
			wg.Done()
		}()
	}

	wg.Wait()

	// sends and receives block until the other side is ready
	// https://stackoverflow.com/questions/24609395/why-is-this-a-deadlock-in-golang-waitgroup
	close(ch)

	for n := range ch {
		println(n)
	}

}

func main2() {
	ch := make(chan int, 3) // buffered
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		// golang也有js的作用域变量共享问题
		j := i

		go func() {
			time.Sleep(time.Second)
			ch <- j
		}()
	}

	go func() {
		for r := range ch {
			println(r)
			wg.Done()
		}
	}()

	wg.Wait()
}
