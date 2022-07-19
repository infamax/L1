package main

import "fmt"

type MySet struct {
	s map[string]struct{}
}

func New() *MySet {
	return &MySet{
		s: make(map[string]struct{}),
	}
}

func (m *MySet) Insert(val string) {
	m.s[val] = struct{}{}
}

func (m *MySet) Exist(val string) bool {
	_, ok := m.s[val]
	return ok
}

func (m *MySet) Erase(val string) {
	_, ok := m.s[val]
	if !ok {
		return
	}
	delete(m.s, val)
}

func (m *MySet) Size() int {
	return len(m.s)
}

func (m *MySet) Print() {
	for key := range m.s {
		fmt.Println(key)
	}
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	m := New()
	for _, val := range strs {
		m.Insert(val)
	}
	m.Print()
}
