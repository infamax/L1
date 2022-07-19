package main

import (
	"log"
	"time"
)

func mySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	log.Println("Start")
	mySleep(5 * time.Second)
	log.Println("finish")
}
