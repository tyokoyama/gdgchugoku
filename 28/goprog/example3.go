package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
channelを2個用意すれば、2個同時に処理させることもできます。
この例では、channelの受信をそれぞれが待つので、どちらが先に終わっても、
両方が終わるまでは次に移ることはありません。
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	c := fanIn(multi("multi(1)"), multi("multi(2)"))
	for i := 0; i < 10; i++ {
		fmt.Printf("Response: %v\n", <-c)
	}
	fmt.Println("main() End")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
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
