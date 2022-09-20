package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func proc(ch chan struct{}) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	go proc2(ch1)
	go proc2(ch2)
	go proc2(ch3)

	select {
	case _, ok := <-ch:
		if !ok {
			// procが増えるとcloseも増える
			// もし3rd libsがこのcloseを忘れてたら？
			close(ch1)
			close(ch2)
			// 試しにコメントアウトしてみると？
			close(ch3)
			log.Println("done from proc")
		}
	}
}

func proc2(ch chan struct{}) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				log.Println("close")
			}
			return
		default:
		}
	}
}

func proc3(ctx context.Context) {
	go proc4(ctx)
	go proc4(ctx)
	go proc4(ctx)
}

func proc4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("done from proc3 err = %s", ctx.Err().Error())
			return
		default:
		}
	}
}

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go proc(ch)
	go proc3(ctx)
	close(ch)
	// context経由で終了を伝播する
	cancel()
	time.Sleep(time.Second * 3)
	log.Println(runtime.NumGoroutine())
}
