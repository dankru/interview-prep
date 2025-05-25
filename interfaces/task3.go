package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	default:
		fmt.Println(v)
	}
}

func task3() {
	do(21)
	do("Hello")
	do(true)
}
