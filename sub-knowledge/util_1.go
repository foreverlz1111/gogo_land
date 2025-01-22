package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
	"unsafe"
)

// defer 用法
func callDefer() {
	defer fmt.Println("world")
	defer fmt.Println("hello")
	// 输出world\n hello
}

// ___
// 哈希表连接
func hashChain() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)
	fmt.Println(s3)
	s4 := append(s1, s3[:2]...)
	fmt.Println(s4)
}

// ___
// Convert 接口的实现为指针
func Convert(value int) []interface{} {
	var slice = make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		slice[i] = value
	}
	return slice
}

// Convert 接口的性能
func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Convert(20)
		_ = result
	}
}

func TestConvert(t *testing.T) {
	result := Convert(20)
	if len(result) != 100 {
		t.Errorf("expected length 100, got %d", len(result))
	}
}

// ___
// Chain 链表结构体
type Chain struct {
	value int
	next  *Chain
}

// 初始化链表
func initChain(val int) *Chain {
	return &Chain{
		val,
		nil,
	}
}

// 获取第n个链表
func accessChain(head *Chain, index int) *Chain {
	for i := 0; i < index; i++ {
		if head == nil {
			return head
		}
		head = head.next
	}
	return head
}

// 查找并打印链表
func findAndPrint(head *Chain, target int) {
	for {
		if head != nil {
			if head.value == target {
				for {
					if head != nil {
						log.Println(head.value)
					} else {
						return
					}
					head = head.next
				}
			}
			head = head.next
		}
	}
}

// ___
func cobraUsage() {
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "A simple CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from CLI")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ___
// 检查逃逸：go build -gcflags="-m" escape_analysis.go
// Go 编译器怎么知道某个变量需要分配在栈上，还是堆上
// 栈内存的分配和释放非常高效。当一个函数被调用时，Go 会为该函数的局部变量分配栈空间，当函数返回时，栈空间会自动释放。放在栈上可以减轻GC压力，减轻内存开销
// 堆: 返回指针、Println、func literal迭代函数、空间不足、使用管道
// 堆是一种相对灵活的内存结构，适用于需要在函数返回后仍然存在的变量，或者大小动态的内存分配
func escapeExample() *int {
	x := 10
	return &x // x 逃逸到堆上，因为它的地址被返回
}

func escapeExample2() int {
	x := 10
	return x // x 不会逃逸，分配在栈上
}
func escapeExample3() {
	name := "string"
	fmt.Println(name) // println 会逃逸
}
func escapeExample4() {
	s := make([]int, 1000, 1000)
	for index := range s {
		s[index] = index // 不会逃逸，分配在栈上
	}
}

// ___

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// ___

// 线程安全的map
type stringMap struct {
	m sync.Map
}

func (s *stringMap) Store(key, value string) {
	s.m.Store(key, value)
}
func (s *stringMap) Load(key string) (value string, ok bool) {
	v, ok := s.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

// ___

// gin框架初始化
var db = make(map[string]string)

func setupGinRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

// ___

// 删除数组的第n个内容
func deleteSliceItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// ___

// 使用waitGroup并发
func waitGroupUsage() {
	c := make(chan int, 3)
	arr := []int{1, 2, 3}

	var wg sync.WaitGroup
	task := func(id int) {
		defer wg.Done() // 在 goroutine 执行完后调用 Done，通知 WaitGroup
		c <- id
	}

	//启动多个 goroutine
	numTasks := 3
	for i := 1; i <= numTasks; i++ {
		wg.Add(1) // 告诉 WaitGroup 要等待一个 goroutine
		go task(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	for range arr {
		u := <-c
		log.Println(u)
	}
	// 所有 goroutine 完成后执行
	fmt.Println("All goroutines finished!")
}

// try运行函数，并尝试恢复
func tryHandler(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// ___

// map用法

func mapUsage() {
	type mymap struct {
		name string
		age  int
	}
	m := make(map[string]mymap)
	m["hello"] = mymap{"a", 1}

	var s []mymap
	for _, v := range m {
		s = append(s, v)
	}

	log.Println(s)
}

func jsonExample() {
	type myJsonExample struct {
		a  string // 无效，字段不是大写
		b  string `json:"B"` // 无效，字段不是大写，有tag也不行
		Cs string // 正常导出为Cs
		D  string `json:"DD"` // 导出为DD,有tag优先tag
	}
	jsonexample := myJsonExample{
		a:  "11",
		b:  "22",
		Cs: "33",
		D:  "44",
	}
	log.Println("myJsonExample:", jsonexample)
	jsonformat, _ := json.Marshal(jsonexample)
	log.Printf("%+v", string(jsonformat))
}

// ___

// go 协程打印字符串，实现并发处理
// goroutine 的运行顺序由 Go 的调度器决定，是非确定性的
// 调度器根据 goroutine 的状态、系统负载和资源可用性来调度执行
// 开发者需要使用同步机制（如 WaitGroup、Mutex、Channels）来管理并发执行和共享数据
// 理解 goroutine 的调度和执行顺序对于编写高效、安全的并发程序至关重要
func goRoutineExample() {
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

// ___

func stringByteUsage() {
	const s = "golangtag"
	println(len(s[:]))
	var a byte = 1 << len(s) / 128    // 常量操作
	var b byte = 1 << len(s[:]) / 512 // 不是常量操作-越位报错
	println(a, b)
}

// ___

// make 是初始化类型为 slice、map、chan 的数据
// make 返回第一个元素,可预设内存空间,避免未来的内存拷贝
// len 和 cap的区别：当前长度与当前容量，append可以进行扩容
// slice 切片是连续内存并且可以动态扩展
func makeUsage() {
	list := make([]int, 4)
	list = append(list, 1, 1)
	fmt.Println(list)

}

// ___

func mapUsage2() {
	m := make(map[int]string)
	m[1] = "hello"
	m[2] = "hi"
	m[3] = "3"
	fmt.Println(m)
	myFuncMap := map[string]func() int{
		"funcA": func() int {
			log.Println("funcA")
			return 1
		},
	}
	f := myFuncMap["funcA"]
	fmt.Println(f())
}

// ___

func syncMapUsage() {
	var m = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			v, _ := m.Load(key)
			fmt.Println(v)
			wg.Done()
		}(i)
		wg.Wait()
	}

	log.Println(m)
}

// ___

// addressUsage
// new 是 生成地址
// *new 返回指针地址
func addressUsage() {
	list := new([]int)
	*list = append(*list, 1)
	// 错误用法：list = append(list, 1)
	fmt.Println(list)
	fmt.Println(*list)
}

// ___

type People interface {
	Speak(string) string
}

type Student struct {
}

func (student Student) Speak(think string) (talk string) {
	println(think)
	talk = "hi"
	return talk
}

// ___

// 平方根计算
func squal() {
	var num int
	num = 913936
	left, right := 0, num
	ctx, cancelctx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelctx()
	for {
		select {
		case <-ctx.Done(): // 检查是否超时
			log.Println("操作超时")
			return
		default:
			mid := (left + right) / 2
			if mid*mid == num { // 找到结果
				log.Println("平方根为:", mid)
				return
			} else if mid*mid > num {
				right = mid
			} else {
				left = mid
			}
			if right-left <= 1 {
				log.Println("未找到精确的平方根，近似值为:", left)
				return
			}
		}
	}
}

// ___

// unsafePointer 指针的不安全用法
func unsafePointer() {
	type Person struct {
		age  int8
		name []string
	}
	person := new(Person)
	log.Println(person)
	// unsafe.Pointer： GO的指针地址不参与运算
	age := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.age))
	name := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.name))
	*((*int)(age)) = 10
	*((*[]string)(name)) = []string{"a", "b"}

	log.Println(person)
}

// ___
func main() {
	//c1 := initChain(1)
	//c2 := initChain(2)
	//c1.next = c2
	//c3 := initChain(3)
	//c2.next = c3
	//c4 := initChain(4)
	//c3.next = c4
	//c5 := initChain(5)
	//c4.next = c5
	//c6 := initChain(6)
	//c5.next = c6
	//log.Println(c1)
	//log.Println(c2)
	//log.Println(c3)
	//findAndPrint(c1, 3)
	// ___

	//fmt.Printf("%d", fib(10))
	// ___

	//stringmap := new(stringMap)
	//log.Println(stringmap)
	// ___

	//r := setupGinRouter()
	//r.Run(":3000") // Listen and Server in 0.0.0.0:8080
	// ___

	//myArray := [5]int{1, 2, 3, 4, 5}
	//fullSlice := myArray[:]
	//remove3rdItem := deleteSliceItem(fullSlice, 2)
	//fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
	// ___

	//tryHandler(func() {
	//	panic("This is panic")
	//}, func(err interface{}) {
	//	log.Println(err)
	//})
	// ___

	//var peo People = Student{}
	//println(peo.Speak("hello"))
	// ___
}
