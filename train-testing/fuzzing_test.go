package main

import (
	"errors"
	"testing"
)

func FuzzCalc(f *testing.F) {
	f.Add(1, 2, "+")
	f.Fuzz(func(t *testing.T, v1, v2 int, ope string) {
		_, _ = Calc(v1, v2, ope)
	})
}

func Calc(v1, v2 int, ope string) (int, error) {
	switch ope {
	case "+":
		return v1 + v2, nil
	case "-":
		return v1 - v2, nil
	case "*":
		return v1 * v2, nil
	case "/":
		return v1 / v2, nil
	}
	return 0, errors.New("")
}
