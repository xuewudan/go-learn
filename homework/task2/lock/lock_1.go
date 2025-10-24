package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int            // 共享的计数器
		mu      sync.Mutex     // 用于保护计数器的互斥锁
		wg      sync.WaitGroup // 用于等待所有协程完成
	)

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，登记一个等待任务
		go func() {
			defer wg.Done() // 协程完成后，标记任务完成

			// 每个协程对计数器进行1000次递增
			for j := 0; j < 1000; j++ {
				mu.Lock()   // 加锁，确保同一时间只有一个协程操作计数器
				counter++   // 递增操作（临界区）
				mu.Unlock() // 解锁，允许其他协程访问
			}
		}()
	}

	wg.Wait() // 等待所有10个协程完成
	fmt.Printf("最终计数器值: %d\n", counter)
}
