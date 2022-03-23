package main

import (
	"fmt"
	"go5/tree"
	"sync"
	"time"
)

func _Say(s string) {
	for i := 1; i <= 3; i++ {
		//等一等
		//time.Sleep(50 * time.Millisecond)
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
func _FibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("退出斐波那契")
			return
		}
	}
}
func Walk(t *tree.Tree, c chan int) {
	//isLRempty = fales
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		Walk(t.Right, c)
	}
}
func SameTree(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	same := false
	Walk(t1, c1)
	Walk(t2, c2)
	for x := 1; x <= len(c1); x++ {
		if <-c1 != <-c2 {
			break
		}
		same = true
	}
	return same
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Increase(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
func main() {

	//go程
	fmt.Println("go程")
	go _Say("World")
	_Say("Hello")
	/**********/

	//go信道
	fmt.Println("go信道")
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	channel1 := make(chan int)
	go _Sum(array1[:len(array1)/2], channel1)
	go _Sum(array1[len(array1)/2:], channel1)
	receive1, receive2 := <-channel1, <-channel1
	fmt.Println("receive1 | receive2 | receive1+receive2 : \n", receive1, receive2, receive1+receive2)
	/**********/

	//go缓冲信道
	fmt.Println("go缓冲信道")
	channel2 := make(chan int, 2)
	channel2 <- 1
	channel2 <- 2
	/*信道溢出程序爆炸
	channel2 <- 3
	*/
	fmt.Println("<- channel2,<- channel2 : ", <-channel2, <-channel2)
	/**********/

	//关闭信道
	fmt.Println("关闭信道")
	channel3 := make(chan int, 10)
	go _Fibonacci(cap(channel3), channel3) //信道的容量
	for v := range channel3 {
		fmt.Printf("_Fibonacci :%v \n", v)
	}
	/**********/

	//select
	fmt.Println("select")
	channel4 := make(chan int)
	channel4_quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel4)
		}
		channel4_quit <- 0
	}()
	_FibonacciSelect(channel4, channel4_quit)
	/**********/

	//select 默认执行
	fmt.Println("select 默认执行")
	tick := time.Tick(50 * time.Millisecond)
	boom := time.After(100 * time.Millisecond)
	isboom := false
	for {
		select {
		case <-tick:
			fmt.Println("计时中...")
		case <-boom:
			fmt.Println("触发！")
			isboom = true
		default:
			fmt.Println("滴...")
			time.Sleep(50 * time.Millisecond)
		}
		if isboom {
			break
		}
	}
	/**********/

	//等价二叉查找树
	fmt.Printf("等价二叉查找树: ")
	tree1 := tree.New(1)
	tree2 := tree.New(2)
	fmt.Println("", SameTree(tree1, tree2))

	//互斥锁
	fmt.Printf("互斥：")
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 90; i++ {
		go c.Increase("xxx")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("xxx"))
}
