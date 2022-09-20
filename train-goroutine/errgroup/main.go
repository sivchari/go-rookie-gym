package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		i := i
		eg.Go(func() error {
			return fmt.Errorf("error i = %d", i)
		})
	}
	// 1つでもerrorがあるとここでキャッチされ、他の処理はcontextを使っているものはcancelされる
	// ただしWithContextを使っていないとcancelされない
	// var eg errgroup.Group
	// またエラーは一番最初の一つしかキャッチされない
	// 複数キャッチはhashicorp multierrorを使ったり自作する
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
	chanerror()
}

func chanerror() {
	ch := make(chan error)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- errors.New("error")
		close(ch)
	}()
	// closeされると抜ける
	for err := range ch {
		if err != nil {
			log.Println(err)
		}
	}
}
