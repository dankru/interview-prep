package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- rand.Intn(100):
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan struct{})
	ch := worker(done)

	for {
		select {
		case num, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(num)
		case <-time.After(time.Second * 5):
			fmt.Println("timeout")
			close(done)
			return
		}
	}
}
