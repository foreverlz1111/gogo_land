package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type vertex struct {
	X int
	Y int
}

var (
	v1 = vertex{2, 3}
	v2 = vertex{X: 1}
	v3 = vertex{}
	v4 = &vertex{1, 2}
)
var array1 [8]int
var str1 [2]string

type mapping struct {
	Lat, Long float64
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
	before, nextt := 0, 1
	return func() int {
		summ := before + nextt
		before = nextt
		nextt = summ
		return before
	}
}

func main() {
	fmt.Println("hello pointer")
	i, j := 53, 443
	p := &i
	fmt.Println("*p = ", *p)
	*p = 21
	fmt.Println("i = ", i)

	p = &j
	*p = *p / 443
	fmt.Println("j = ", j)

	fmt.Println("struct :", vertex{3, 7})

	vert := vertex{3, 7}
	vert.X = 5
	vv := &vert
	vv.X = 1e9
	fmt.Println("v.X = ", vert.X)
	fmt.Println("vv.x= ", vv.X)

	fmt.Println(v1, v2, v3, v4)

	str1[0] = "Hello"
	str1[1] = "GOlang"
	for a := 0; a < len(str1); a++ {
		fmt.Println("str1[", a, "]= ", str1[a])
	}
	fmt.Println("array1 int len() = ", len(array1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for a := 0; a < len(array1); a++ {
		array1[a] = r.Intn(66)
	}
	array2 := array1[4:7] //array2[4]、array2[5]、array2[6]这三个
	fmt.Println("array1 = ", array1)
	fmt.Println("array2 = ", array2)

	names := [4]string{"name1", "name2", "name3", "name4"} //[4]中的数字4可以省略
	fmt.Println(names)
	name_b := names[1:3]
	fmt.Println(names[0:2], name_b)

	array3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("array3 = ", array3[0:3])
	array4 := []bool{true, true, false}
	fmt.Println("array4", array4)
	array5 := []struct {
		a int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	} //严格的缩进
	fmt.Println("array5", array5)

	printSlice(array3)

	s := array3[:0] //取0个
	printSlice(s)
	s = s[:4] //取四个
	printSlice(s)
	s = s[2:] //舍弃前两个
	printSlice(s)

	//s = s[6:6]//此时s也是空len()和空cap()但不显示为nil

	var array_emp []int
	fmt.Println(array_emp, len(array_emp), cap(array_emp))
	if array_emp == nil {
		fmt.Println("空数组！")
	}

	array6 := make([]int, 5) //make 函数分配一个元素为零的数组
	fmt.Printf("array6 ")
	printSlice(array6)

	array7 := make([]int, 0, 5) //make(array,len,cap)
	fmt.Printf("array7 ")
	printSlice(array7)

	array8 := array7[:2]
	fmt.Printf("array8 ")
	printSlice(array8)

	array9 := array8[2:5]
	fmt.Printf("array9 ")
	printSlice(array9)

	array10 := [][]string{
		make([]string, 3),
		make([]string, 3),
		make([]string, 3),
	}
	for i := 0; i < len(array10); i++ {
		for j := 0; j < len(array10); j++ {
			if i == 0 {
				array10[i][j] = "Y"
			} else {
				array10[i][j] = "K"
			}
		}
	}

	for i := 0; i < len(array10); i++ {
		for j := 0; j < len(array10); j++ {
			fmt.Printf("%s \n", strings.Join(array10[i], " "))
		}
	}
	array11 := []int{2, 4, 5, 7, 3, 1, 4}
	for i, v := range array11 {
		fmt.Printf("index(%d) = %d \n", i, v) //下标+下标对应的值
	}
	array12 := make([]int, 10)
	for i := range array12 {
		array12[i] = 1 << uint(i) // 2^i
	}
	for _, value := range array12 {
		//赋予下标 _ 进行忽略
		fmt.Printf("array12[] = %d \n", value)
	}
	map1 := make(map[string]mapping) //必须用make初始化后使用
	map1["str1"] = mapping{
		1, 2,
	}
	map2 := map[string]mapping{ //也可以不使用make初始化使用
		"str22": mapping{
			11, 22,
		},
		"str33": mapping{
			111, 222,
		},
	}
	map3 := map[string]mapping{ //更简洁的定义
		"str44": {1111, 2222},
		"str55": {11111, 22222},
	}

	fmt.Println("map1 映射：", map1)
	fmt.Println("map2 映射：", map2)
	fmt.Println("map3 映射：", map3)

	tmp := mapping{
		Lat: 1,
	}
	map3["str44"] = tmp
	fmt.Println("map3 修改后的映射：", map3)

	fmt.Printf("Filed : %q \n", strings.Fields("23 hhh string"))

	func_num := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("func_num :", func_num)
	fmt.Println("compute(func_num) :", compute(func_num))
	fmt.Println("compute(math.Pow)", compute(math.Pow))

	pos_num, neg_num := adder(), adder()
	fmt.Printf("闭包：")
	for i := 0; i < 5; i++ {
		fmt.Println(pos_num(i), neg_num(-2*i))
	}

	fib := fibonacci()
	for i := 0; i < 14; i++ {
		fmt.Printf("斐波那契数列 第%d次: %d\n", i, fib())
	}
}
func printSlice(s []int) {
	fmt.Printf("长度= %d ，首个元素 = %d \"array\"= %v\n", len(s), cap(s), s)
}
