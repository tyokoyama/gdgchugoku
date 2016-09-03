package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   呼び出し側から終了させるためのchannelを渡し、
   呼び出し側が終了するときにchannelに値を送信する。
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	quit := make(chan bool)
	c := multi("mutli", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- true

	<-quit
	fmt.Println("main() End.")
}

func multi(st string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", st, i):
			case <-quit:
				// 何か後処理をしてから、呼び出し側に返事を返す。
				cleanup()
				quit <- true
				return
			}
		}
	}()

	return c
}

func cleanup() {
	// 後処理
	fmt.Println("cleanup()")
}
