package main

import (
	"fmt"
	"testing"
)

func BenchmarkAppendAllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		// 都度append
		base = append(base, fmt.Sprintf("No %d", i))
	}
}

func BenchmarkAppendAllocateOnce(b *testing.B) {
	//最初に長さを決める
	base := make([]string, b.N)
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("No %d", i)
	}
}
