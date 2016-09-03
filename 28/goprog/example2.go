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

	c1 := multi("multi(1)")
    c2 := multi("multi(2)")
	for i := 0; i < 5; i++ {
		fmt.Printf("Response(1): %v\n", <-c1)
		fmt.Printf("Response(2): %v\n", <-c2)
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
