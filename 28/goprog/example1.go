package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
channelを使って、goroutineからの応答を待つ。
今回の場合、buffer channelではないので、1つ値が送信されるまで
受信側はずっと待たされる。
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	c := multi("multi")
	for i := 0; i < 5; i++ {
		fmt.Printf("Response: %v\n", <-c)
	}
	fmt.Println("main() End")
}

func multi(st string) <-chan string {
	fmt.Printf("multi is called.\n")
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", st, i)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
	}()
	fmt.Printf("return\n")
	return c
}
