package main

import (
	"log"
	"sync"
	"time"
)

type mem struct {
	mem map[int]int
	mu  sync.Mutex
}

func main() {
	m := mem{
		mem: make(map[int]int),
	}
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			m.mu.Lock()
			m.mem[i] = i
			m.mu.Unlock()
		}()
	}

	time.Sleep(time.Second * 3)
	for k, v := range m.mem {
		log.Println(k, v)
	}
	log.Println("end")
}
