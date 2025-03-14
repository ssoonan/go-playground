package main

import (
	"fmt"
)

// type Speaker interface {
// 	Speak() string
// }

type Person struct {
	Name string
}

func (p Person) Speak () string {
	return "Hello, I'm " + p.Name
}

func (p* Person) Speak2 () string {
	return "Hello, I'm " + p.Name
}

func main() {
	var s Person
	p := Person{Name: "Alice"}
	s = p                  // 값 타입(Person)으로 인터페이스 만족
	fmt.Println(s.Speak()) // "Hello, I'm Alice"

	s = &p                 // 포인터 타입(*Person)으로도 인터페이스 만족
	fmt.Println(s.Speak2()) // "Hello, I'm Alice"
}
