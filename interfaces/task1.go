package main

import "fmt"

// task1
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func task1() {
	var s Speaker
	fmt.Println(s == nil) // выведет true потому что оба поля интерфейса nil
	s = Dog{}
	fmt.Println(s == nil) // выведет false потому что поле itab интерфейса = Dog
}
