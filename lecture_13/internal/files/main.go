package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	mx := &sync.RWMutex{}

	go write(m, mx)
	go read(m, mx)
	go read(m, mx)

	block := make(chan struct{})
	<-block
}

func write(m map[int]int, mx *sync.RWMutex) {
	for {
		for i := 0; i < 10; i++ {
			mx.Lock()
			m[i] = i
			mx.Unlock()
		}
	}
}

func read(m map[int]int, mx *sync.RWMutex) {
	for {
		mx.RLock()
		for k, v := range m {
			fmt.Println(k, " - ", v)
		}
		mx.RUnlock()
	}
}
