package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// HTTP Requestのtimeoutを決めたい
	// ある処理に指定秒以上をかけられない
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	go watch(ctx)
	log.Println("execute ...")
	time.Sleep(time.Second * 5)
	// cancel()
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		default:
			time.Sleep(time.Second * 3)
			log.Println("watch ...")
		}
	}
}
