package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Add(1, 1))
	fmt.Println(Add(1, -1))
}

func Add(i, j int) (int, error) {
	if i < 0 || j < 0 {
		return 0, errors.New("invalid")
	}
	return i + j, nil
}
