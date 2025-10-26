package one

import (
	"fmt"
)

func ChannelTest() {
	fmt.Println("main start")
	ch := make(chan string)
	ch <- "a" // 入 chan
	go func() {
		val := <-ch // 出 chan
		fmt.Println(val)
	}()
	fmt.Println("main end")
}
