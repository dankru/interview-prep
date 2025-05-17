package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer fmt.Println("time took for execution: ", time.Since(t))

	bufChan := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			bufChan <- i
		}
		close(bufChan)
	}()

	for num := range bufChan {
		fmt.Println(num)
	}
}
