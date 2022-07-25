package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action
	от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) Name() string {
	return fmt.Sprintf("My name is %s", h.name)
}

type Action struct {
	Human
}

func (a *Action) Eat() {
	fmt.Println("I am eating")
}

func main() {
	var a Action
	a.SetName("Maxim")
	fmt.Println(a.Name())
	a.Eat()
}
