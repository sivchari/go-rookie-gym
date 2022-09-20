package main

import (
	"log"
	"sync"
)

func main() {
	// waitgroupは内部のカウンタを用いるためfuncに渡すときはコピーではなく参照を渡す
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		i := i
		// カウントを増やす
		wg.Add(1)
		go func(wg sync.WaitGroup) {
			// カウントを減らす
			defer wg.Done()
			log.Println(i)
		}(wg)
	}
	// カウントが0でない限りブロックする
	wg.Wait()
	log.Println("end")
}
