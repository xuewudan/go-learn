package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个带缓冲的通道，缓冲区大小为10（可根据需求调整）
	bufferedChan := make(chan int, 10)
	var wg sync.WaitGroup

	// 生产者协程：发送1-100的整数到通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			// 向缓冲通道发送数据
			// 当缓冲区未满时直接发送，满了则阻塞等待消费者取走数据
			bufferedChan <- i
			// 可选：打印发送信息（便于观察缓冲区工作状态）
			// fmt.Printf("发送: %d (当前缓冲区长度: %d)\n", i, len(bufferedChan))
		}
		// 所有数据发送完成后关闭通道
		close(bufferedChan)
		fmt.Println("生产者：所有数据发送完毕，关闭通道")
	}()

	// 消费者协程：从通道接收数据并打印
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 循环接收数据，直到通道关闭且缓冲区为空
		for num := range bufferedChan {
			fmt.Printf("接收: %d\n", num)
			// 模拟处理耗时（可选，用于观察缓冲区变化）
			// time.Sleep(10 * time.Millisecond)
		}
		fmt.Println("消费者：所有数据接收完毕")
	}()

	// 等待生产者和消费者完成
	wg.Wait()
	fmt.Println("程序结束")
}
