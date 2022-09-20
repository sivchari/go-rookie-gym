package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 最大実行数は1
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// Acquireで取得するとcountが減る
	// 0になるとブロックする
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	// カウントを戻す
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
