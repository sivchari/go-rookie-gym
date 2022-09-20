package main

import (
	"errors"
	"log"
)

func main() {
	ch := make(chan int)
	errch := make(chan error)
	go func() {
		for i := 0; i < 10; i++ {
			if i != 9 {
				ch <- i
			} else {
				errch <- errors.New("error")
			}
		}
	}()

LOOP:
	for {
		select {
		case i := <-ch:
			log.Println(i)
		case err := <-errch:
			log.Println(err)
			break LOOP
		}
	}
	log.Println("end")
}
