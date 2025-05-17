package main

import "fmt"

func processor(done <-chan struct{}, nums <-chan int, squares chan<- int, cubes chan<- int) {
	defer close(squares)
	defer close(cubes)
	for {
		select {
		case <-done:
			return
		case num := <-nums:
			squares <- num * num
			cubes <- num * num * num
		}
	}
}

func main() {
	done := make(chan struct{})
	nums := make(chan int)
	squares := make(chan int)
	cubes := make(chan int)

	go processor(done, nums, squares, cubes)

	go func() {
		for i := 0; i < 5; i++ {
			nums <- i
		}
		close(nums)
		close(done)
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("squares: %d, cubes: %d\n", <-squares, <-cubes)
	}
}
