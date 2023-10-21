package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cnt := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			for j := 0; j < 10; j++ {
				cnt += 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
