package main

import (
	"context"
	"log"
	"unsafe"
)

type keytype struct{}

var key keytype = struct{}{}

func main() {
	ctx := context.WithValue(context.Background(), "key", 1)
	log.Println(ctx.Value("key"))

	ctx2 := context.WithValue(context.Background(), key, 1)
	log.Println(ctx2.Value(key))

	log.Println(unsafe.Sizeof(key))
	log.Println(unsafe.Sizeof("key"))
}
