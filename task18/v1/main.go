package main

import (
	"fmt"
	"sync"
	"time"
)

// Первый способ используя mutex
type counter struct {
	count int
	mu    *sync.RWMutex
}

func New(count int) *counter {
	return &counter{
		count: count,
		mu:    &sync.RWMutex{},
	}
}

func (c *counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *counter) Value() int {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.count
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	c := New(0)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value())
	fmt.Println("With mutex:", time.Now().Sub(start).Seconds())
}
