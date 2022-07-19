package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Используя пакет атомик
type counter struct {
	count int64
}

func New(count int64) *counter {
	return &counter{
		count: count,
	}
}

func (c *counter) Inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.count)
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
	fmt.Println("value :", c.Value())
	fmt.Println("with atomic: ", time.Now().Sub(start).Seconds())
}
