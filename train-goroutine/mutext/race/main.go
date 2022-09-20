package main

import (
	"fmt"
	"time"
)

func main() {
	c := 0
	for i := 0; i < 1000; i++ {
		go func() {
			c++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c)
}
