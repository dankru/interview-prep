package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Считывает уведомления из воркера, пишет в 2 других канала
type DataProvider interface {
	getData() <-chan int
	stop()
}

type Notifier struct {
}

func (n *Notifier) getNotification(w Worker) []int {
	nums := make([]int, 0, 10)
	defer w.stop()

	for v := range w.dataCh {
		nums = append(nums, v)
		if len(nums) > 1 {
			break
		}
	}

	return nums
}

type Worker struct {
	dataCh  chan int
	closeCh chan struct{}
}

func getData() Worker {
	w := Worker{
		dataCh:  make(chan int),
		closeCh: make(chan struct{}),
	}

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				w.dataCh <- rand.Intn(100)
			case <-w.closeCh:
				return
			}
		}
	}()

	return w
}

func (w *Worker) stop() {
	w.closeCh <- struct{}{}
	close(w.closeCh)
}

func main() {
	n := Notifier{}

	w := getData()
	nums := n.getNotification(w)
	fmt.Println(nums)
}
