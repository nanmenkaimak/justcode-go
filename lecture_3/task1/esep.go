package main

import (
	"fmt"
	"sync"
)

func firstExample() {
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}

	go func() {
		mu1.Lock()
		mu2.Lock()
		defer mu1.Unlock()
		defer mu2.Unlock()
	}()

	go func() {
		mu2.Lock()
		mu1.Lock()
		defer mu2.Unlock()
		defer mu1.Unlock()
	}()

	fmt.Println("Waiting for goroutines to finish...")
	select {}
}

func secondExample() {
	var i int
	num := make(chan int)
	i = 1
	num <- i
	num <- i + 2
	fmt.Println(num)
}
