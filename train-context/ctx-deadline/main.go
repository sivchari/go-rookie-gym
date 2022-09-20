package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
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
