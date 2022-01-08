package main

import (
	"fmt"
	"time"
	"tour/tree"
)

func _Say(s string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i, ":", s)
	}
}
func _Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func _Fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func _Fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("退出！")
			return
		}
	}
}
func main() {
	//go程
	go _Say("world")
	_Say("")
	/**********/

	//go信道
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	channel1 := make(chan int)
	go _Sum(array1[:len(array1)/2], channel1)
	go _Sum(array1[len(array1)/2:], channel1)
	receive1, receive2 := <-channel1, <-channel1
	fmt.Println("receive1 | receive2 | receive1+receive2 : \n", receive1, receive2, receive1+receive2)
	/**********/

	//go缓冲信道
	channel2 := make(chan int, 2)
	channel2 <- 1
	channel2 <- 2
	/*信道溢出程序爆炸
	channel2 <- 3
	*/
	fmt.Println("<- channel2,<- channel2 : ", <-channel2, <-channel2)
	/**********/

	//关闭信道
	channel3 := make(chan int, 10)
	go _Fibonacci(cap(channel3), channel3) //信道的容量
	for v := range channel3 {
		fmt.Printf("_Fibonacci :%v \n", v)
	}
	/**********/

	//select
	channel4 := make(chan int)
	channel4_quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel4)
		}
		channel4_quit <- 0
	}()
	_Fibonacci_select(channel4, channel4_quit)
	/**********/

	//select 默认执行
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for{
		select{
			case <- tick:
			fmt.Println("计时中...")
			case <- boom:
			fmt.Println("触发！")
			return
			default:
			fmt.Println("默认情况")
			time.Sleep(50*time.Millisecond)
		}
	}
	/**********/

	//等价二叉查找树
	my_tree := tree.Tree{}
}
