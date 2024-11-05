package main

import (
	"fmt"
	"sync"
	"time"
)

// go 协程打印字符串，实现并发处理
// goroutine 的运行顺序由 Go 的调度器决定，是非确定性的
// 调度器根据 goroutine 的状态、系统负载和资源可用性来调度执行
// 开发者需要使用同步机制（如 WaitGroup、Mutex、Channels）来管理并发执行和共享数据
// 理解 goroutine 的调度和执行顺序对于编写高效、安全的并发程序至关重要
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	str := "Hello World"
	str1 := []byte(str)
	sc := make(chan byte, len(str))
	count := make(chan int)
	for _, v := range str1 {
		sc <- v
	}
	close(sc)

	go func() {
		defer wg.Done()
		for {
			ball, ok := <-count
			if ok {
				pri, ok1 := <-sc
				if ok1 {
					fmt.Printf("go 1:%c \n", pri)
				} else {
					close(count)
					return
				}
				count <- ball
			} else {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(2 * time.Millisecond)
			ball, ok := <-count
			if ok {
				pri, ok1 := <-sc
				if ok1 {
					fmt.Printf("go 2: %c \n", pri)
				} else {
					close(count)
					return
				}
			} else {
				return
			}
			count <- ball
		}
	}()
	count <- 23
	wg.Wait()
}
