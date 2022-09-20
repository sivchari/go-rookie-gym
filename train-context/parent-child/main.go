package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", 1)
	log.Println("-----")
	log.Println(ctx.Value("key"))
	log.Println(ctx.Value("key2"))

	log.Println("-----")
	ctx2 := context.WithValue(ctx, "key2", 2)
	log.Println(ctx2.Value("key2"))

	log.Println("-----")
	ctx3 := context.WithValue(ctx2, "key3", 3)
	log.Println(ctx3.Value("key2"))

	log.Println("-----")
	ctx4 := context.WithValue(ctx, "key4", 4)
	log.Println(ctx4.Value("key2"))

	ctxtimeout, _ := context.WithTimeout(context.Background(), time.Second*3)
	// timeoutが短い方が優先される
	childctxtimeout, _ := context.WithTimeout(ctxtimeout, time.Second*2)
	child2ctxtimeout, _ := context.WithTimeout(ctxtimeout, time.Second*4)
	time.Sleep(time.Millisecond * 2500)
	log.Println("-----")
	log.Println(childctxtimeout.Err()) // deadline
	log.Println(ctxtimeout.Err())
	time.Sleep(time.Millisecond * 1000)
	log.Println(child2ctxtimeout.Err()) // deadline 4秒でなく、3.5秒でcancelされてる

	parent, cancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	cancel()
	// 親から子へはcancelが伝播する
	log.Println("-----")
	log.Println(parent.Err())
	log.Println(child.Err())

	parent2, _ := context.WithCancel(context.Background())
	child2, cancel := context.WithCancel(parent)
	cancel()
	// 子から親へは伝播しない
	log.Println("-----")
	log.Println(parent2.Err())
	log.Println(child2.Err())
}
