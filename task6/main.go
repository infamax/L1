package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// StopGoroutinesV1 Первый способ используя контекст
func StopGoroutinesV1() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// Some work here
	}()
	wg.Wait()
}

// StopGoroutinesV2 Используя канал
func StopGoroutinesV2() {
	exit := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-exit
		// Some work here
	}()
	time.Sleep(2 * time.Second)
	exit <- true
	wg.Wait()
}

// StopGoroutinesV3 остановка с помощью сигнала
func StopGoroutinesV3() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
	}()
	wg.Wait()
}

func main() {
	log.Println("start first")
	StopGoroutinesV1()
	log.Println("finish first")
	log.Println("start second")
	StopGoroutinesV2()
	log.Println("finish second")
	log.Println("start third")
	log.Println("Enter ctrl + c")
	StopGoroutinesV3()
	log.Println("finish third")
}
