package main

import (
	"fmt"
	"sync"
)

// 打印1-10的奇数
func printOdds(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知WaitGroup
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

// 打印2-10的偶数
func printEvens(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知WaitGroup
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

func main() {
	var wg sync.WaitGroup

	// 注册2个需要等待的协程
	wg.Add(2)

	// 启动协程打印奇数
	go printOdds(&wg)
	// 启动协程打印偶数
	go printEvens(&wg)

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有打印完成")
}
