package main

import (
	"sync"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	racemem := make(map[int]int)
	go func() {
		racemem[0] = 0
	}()
	racemem[1] = 1
	time.Sleep(time.Second * 3)
}

func TestNoRace(t *testing.T) {
	type noracemem struct {
		mu  sync.Mutex
		mem map[int]int
	}
	nmem := noracemem{
		mem: make(map[int]int),
	}
	go func() {
		nmem.mu.Lock()
		defer nmem.mu.Unlock()
		nmem.mem[0] = 0
	}()
	nmem.mu.Lock()
	nmem.mem[1] = 1
	nmem.mu.Unlock()
	time.Sleep(time.Second * 3)
}
