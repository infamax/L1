package main

import (
	"fmt"
	"github.com/infamax/L1/task24/point"
)

func main() {
	p1 := point.New(0, 0)
	p2 := point.New(3, 4)
	fmt.Println(point.Distance(p1, p2))
}
