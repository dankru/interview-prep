package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func task2() {
	var i interface{}
	describe(i) // Выведет type = nil value = nil

	i = 42
	describe(i) // Выведет type = int value = 42

	i = "hello"
	describe(i) // Выведет type = string value = hello
}
