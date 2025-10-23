package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	oddch := make(chan int, 5)
	evench := make(chan int, 5)
	oddquit := make(chan int, 1)
	evenquit := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go odd(oddch, oddquit, &wg)
	go even(evench, evenquit, &wg)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				evench <- i
			} else {
				oddch <- i
			}
		}
		close(oddch)
		close(evench)
	}()

	// Wait for the sender goroutine to finish
	time.Sleep(100 * time.Millisecond)

	oddquit <- 0
	evenquit <- 0

	wg.Wait()
}

func odd(ch <-chan int, quit <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("odd", v)
		case <-quit:
			fmt.Println("odd quit")
			return
		}
	}
}

func even(ch <-chan int, quit <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("even", v)
		case <-quit:
			fmt.Println("even quit")
			return
		}
	}
}
