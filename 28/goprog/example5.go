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
	c := multi("multi", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- true
}

func multi(st string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ;i++ {
			select {
			case c <- fmt.Sprintf("%s %d", st, i):
			case <-quit:
				fmt.Println("Quit")
				return
			}
		}
	}()

	return c
}
