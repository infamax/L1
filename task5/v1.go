package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func GettingRandomNumbers(ctx context.Context, N int) {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(N)*time.Second)
	defer cancel()
	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				log.Printf("Sending number: %d", i)
				ch <- i
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for val := range ch {
		log.Printf("Getting number from channel: %d", val)
	}
}

func main() {
	var N int
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start")

	ctx := context.Background()
	GettingRandomNumbers(ctx, N)
	log.Println("Finish")
}
