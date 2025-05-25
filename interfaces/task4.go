package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyType struct {
	Value string
}

func (m *MyType) Method() {
	if m == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("Not nil: ", m.Value)
}

func task4() {
	var i MyInterface // nil т.к. itab и data nil
	var t *MyType     // nil т.к. zerovalue указателя

	i = t      // уже не nil, т.к. itab = *MyType, data по прежнему nil
	i.Method() // выведет <nil> т.к. сам интерфейс не nil из-за itab, но в метод передаётся не itab, а data. В аргумент m *MyType прилетает содержимое data

	i = &MyType{"Hello"} // не nil, т.к. itab = *MyType, data = указатель на структуру
	i.Method()           // выведет Hello
}
