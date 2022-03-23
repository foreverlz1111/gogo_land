package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
	//包导入但未使用=报错
)

var c, python, java bool //函数外必须用var定义
var i, j int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	//var input_value int
	//fmt.Scanf("%d",&input_value)//使用fmt.Scan(&input_value)不指定输入类型
	fmt.Println("shining in  go\n", time.Now())
	fmt.Println("现在是随机数:(固定不变的)", rand.Intn(199))
	fmt.Println("9开根：", math.Sqrt(9))
	fmt.Println("Pi :", math.Pi)
	fmt.Println("add(7,8)= ", add(7, 8))
	hello, world := swap("hello", "world")
	fmt.Println("hello=,world=", hello, world)
	fmt.Println(split(12))
	var number int
	fmt.Println("number,c,python,java =", number, c, python, java)
	var c, python, java = true, false, "no"
	fmt.Println("i,j,c,python,java = ", i, j, c, python, java)
	var i, j int = 1, 2
	k := 3 //只有新变量才使用 := 进行简洁定义
	fmt.Println("i,j,k=", i, j, k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z) //复数
	var f float64
	var b bool
	var s string
	fmt.Printf("f, b, s = %v %v %q\n", f, b, s)
	x, y := 3, 4
	f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f) //不同格式变量必须使用显式转换
	fmt.Println("x,y,z=", x, y, z)
	v := 42 // 变量类型取决于赋的值
	fmt.Printf("v is of type： %T\n", v)
	const World = "世界"
	fmt.Println("Hello", World)
	//Pi := 3.1
	fmt.Println("Happy", Pi, "Day") //静态常量用const定义，且不能用=赋值.使用:=会重新定义先前声明的静态常量使之成为变量
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println("needInt(Small)", needInt(Small))
	fmt.Println("needFloat(Small)", needFloat(Small))
	fmt.Println("needFloat(Big)", needFloat(Big))
}
func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 5 / 6
	y = sum - x
	return
}
func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
