package main

import (
	"time"
)

func main() {
	ch := make(chan struct{})
	for i := 0; i < 5; i++ {
		goroutine(ch)
	}
	time.Sleep(1 * time.Second)
	close(ch)
}

func goroutine(ch chan struct{}) {
	go func() {
		var counter int64
		for {
			select {
			case <-ch:
				return
			default:
				counter++
			}
		}
	}()
}
