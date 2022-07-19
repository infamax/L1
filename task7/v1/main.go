package main

import (
	"fmt"
	"log"
	"sync"
)

// Реализовать конкурентную запись данных в map.
// Первая реализация используя mutex
type concurrentMap struct {
	m  map[int]int
	mu *sync.RWMutex
}

func New() *concurrentMap {
	return &concurrentMap{
		m:  map[int]int{},
		mu: &sync.RWMutex{},
	}
}

func (c *concurrentMap) Add(key, val int) {
	c.mu.Lock()
	c.m[key] = val * 2
	c.mu.Unlock()
}

func (c *concurrentMap) Value(key int) int {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.m[key]
}

func (c *concurrentMap) Erase(key int) {
	defer c.mu.Unlock()
	c.mu.Lock()
	_, ok := c.m[key]
	if !ok {
		return
	}
	delete(c.m, key)
}

func main() {
	cm := New()
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Printf("Recording key %d, value %d\n", i, i*i)
			cm.Add(i, i*i)
			fmt.Printf("key = %d, value = %d\n", i, cm.Value(i))
		}(i)
	}
	wg.Wait()

}
