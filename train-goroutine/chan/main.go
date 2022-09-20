package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()
	log.Println("wait")
	log.Println(<-ch)
	close(ch)
	go func() {
		time.Sleep(time.Second * 2)
		// closeしても遅れるが初期値が返却されるようになる
		ch <- 1
	}()
	log.Println(<-ch)
	// closeを2回するとpanicする
	// close(ch)
	log.Println("end")
}
