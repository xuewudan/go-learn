package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// 发送协程
	wg.Add(1) // 登记一个待完成的协程
	go func() {
		defer wg.Done() // 协程完成后标记
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	// 接收协程
	wg.Add(1) // 再登记一个待完成的协程
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收: %d\n", num)
		}
	}()

	wg.Wait() // 等待所有登记的协程完成
	fmt.Println("程序结束")
}
