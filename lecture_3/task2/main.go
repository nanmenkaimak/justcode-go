package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 4
		close(ch1)
		ch2 <- 2
		ch2 <- 56
		close(ch2)
		ch3 <- 3
		ch3 <- 6688
		close(ch3)
	}()

	for num := range merge(ch1, ch2, ch3) {
		fmt.Println(num)
	}
}

func merge(chans ...chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, n := range chans {
		go func(ch chan int) {
			defer wg.Done()
			for num := range ch {
				res <- num
			}
		}(n)
	}

	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}
