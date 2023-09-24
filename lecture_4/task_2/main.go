package main

import (
	"fmt"
	"sync"
)

type customMap struct {
	m  map[string]int
	mu sync.Mutex
}

func newMap() *customMap {
	return &customMap{
		m: make(map[string]int),
	}
}

func (c *customMap) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[key]
}

func (c *customMap) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func main() {
	myMap := newMap()
	myMap.Set("first", 1)
	fmt.Println(myMap.Get("first"))
}
