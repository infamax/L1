package main

import (
	"fmt"
	"sync"
	"time"
)

// Используя каналы
type counter struct {
	count      int
	countChan  chan int
	addingChan chan int
}

func New(count int) *counter {
	c := &counter{
		count:      count,
		countChan:  make(chan int),
		addingChan: make(chan int),
	}
	go c.run()
	return c
}

func (c *counter) Value() int {
	return <-c.countChan
}

func (c *counter) Inc(amount int) {
	c.addingChan <- amount
}

func (c *counter) applyDelta(amount int) {
	c.count += amount
}

func (c *counter) run() {
	var delta int
	for {
		select {
		case delta = <-c.addingChan:
			c.applyDelta(delta)
		case c.countChan <- c.count:
			//
		}
	}
}

func main() {
	start := time.Now()
	c := New(0)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc(1)
		}()
	}
	wg.Wait()
	fmt.Println("value ", c.Value())
	fmt.Println("With channels:", time.Now().Sub(start).Seconds())
}
