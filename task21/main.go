package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

// Есть интерфейс USB с методом ConnectWithUSBCable()
// и структура реализующая карту памяти MemoryCard
// С помощью адаптера (CardReader) мы делаем возможным прочитать данные с карты

type USB interface {
	ConnectWithUSBCable()
}

type MemoryCard struct {
}

func (m *MemoryCard) Insert() {
	fmt.Println("Карта памяти успешна вставлена!")
}

func (m *MemoryCard) CopyData() {
	fmt.Println("Данный успешно скопированы!")
}

type CardReader struct {
	m *MemoryCard
}

func (c *CardReader) ConnectWithUSBCable() {
	c.m.Insert()
	c.m.CopyData()
}

func main() {
	c := CardReader{}
	c.ConnectWithUSBCable()
}
