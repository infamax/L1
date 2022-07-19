package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func WriteRandomDataInChan(ch chan<- int, osChan chan os.Signal) {
	for {
		select {
		case <-osChan: // Поступил сигнал завершить работу
			log.Printf("main goroutine stop the work\n")
			return
		case <-time.After(1 * time.Second): // Раз в секунду пишем рандомные данные в канал
			rand.Seed(time.Now().UnixNano())
			num := rand.Int()
			log.Printf("recording main goroutine value %d\n", num)
			ch <- num
		}
	}
}

func StartWorkers(ctx context.Context, wg *sync.WaitGroup, countWorkers int, ch <-chan int) {
	wg.Add(countWorkers)
	for i := 0; i < countWorkers; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select { // Завершаем работу горутины
				case <-ctx.Done():
					log.Printf("end working worker %d", i)
					return
				case inputData := <-ch: // Получаем данные из главного канала
					log.Printf("worker %d, get data from main channel %d", i, inputData)
				}
			}
		}(i)
	}
}

func main() {
	var countWorkers int
	_, err := fmt.Scanf("%d", &countWorkers)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("start")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	osChan := make(chan os.Signal)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)
	ch := make(chan int)
	wg := sync.WaitGroup{}
	StartWorkers(ctx, &wg, countWorkers, ch)
	WriteRandomDataInChan(ch, osChan)
	cancel()
	wg.Wait()
	log.Println("finish")
}
