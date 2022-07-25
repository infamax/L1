package main

import (
	"log"
	"time"
)

// Реализовать собственную функцию sleep.

func mySleep1(d time.Duration) {
	<-time.After(d)
}

func mySleep2(ns int64) {
	if ns <= 0 {
		return
	}

}

func main() {
	log.Println("Start")
	mySleep1(5 * time.Second)
	log.Println("finish")
}
