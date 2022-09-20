package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Printf("before leak:%d\n", runtime.NumGoroutine())

	leak(nil)

	time.Sleep(3 * time.Second)
	log.Printf("after leak:%d\n", runtime.NumGoroutine())
}

func leak(c <-chan string) {
	// closeされない！！
	go func() {
		for cc := range c {
			log.Println(cc)
		}
	}()

	log.Printf("in leak:%d\n", runtime.NumGoroutine())
}
