package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
	"os"
	"image"
	"image/color"
	"tour/pic"
)

type vertex1 struct {
	X, Y float64
}
type myfloat float64

//****************************************
func (v *vertex1) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* 非指针：
   func abs(v vertex1)float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
        }
*****************************************/
func (v *vertex1) scale(f float64) {
	v.X = v.Y * f
	v.Y = v.X * f
}

//****************************************
func scale(v *vertex1, f float64) {
	v.X = v.Y * f
	v.Y = v.X * f
}

/* 非指针：
   func scale(v vertex1,f float64){
	v.X = v.Y * f
	v.Y = v.X * f
        }
*****************************************/
func (myf myfloat) abs() float64 {
	if myf < 0 {
		return float64(-myf)
	}
	return float64(myf)
}

type abser interface {
	abs() float64
}

//********************1接口——2结构——3方法
type interf interface {
	matt()
}
type typp struct {
	str string
}

func (t *typp) matt() {
	if t == nil {
		fmt.Println("t的值为 null!")
		return
	}
	fmt.Println(t.str)
	/**********非指针
	  func (t typp) matt(){
	  fmt.Println(t.str)
	  }
	  **********/
} /********************/
func (f F) matt() {
	fmt.Println("f = ", f)
}

type F float64

func describe(i interf) {
	fmt.Printf("(数值：%v,类型： %T)\n", i, i)
}
func _whattype(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf(" %v ,乘以二 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 类型是 , 共 %v 字节\n", v, len(v))
	default:
		fmt.Printf("未定义识别到的类型，(%T) \n", v)
	}
}

type person struct {
	_name string
	_age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years old)", p._name, p._age)
}

type _ipaddress [4]byte

//***********1错误结构定义，2错误封装，3函数错误定义
type _error_output struct {
	when time.Time
	what string
}

func (e *_error_output) Error() string {
	//封装错误信息
	return fmt.Sprintf("时间 : %v\n信息 : %s", e.when, e.what)
}
func _Run() error {
	return &_error_output{time.Now(), "错误"}
}

/***********/
//**********封装一个开方错误
type _sqrt_value float64

func (x _sqrt_value) Error() string {
	//在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环
	return fmt.Sprintf("输入值:%f 信息:输入负数！", x)
}
func _Sqrt(x _sqrt_value) (float64, error) {
	if x > 0 {
		//math.Sqrt()接受float64类型
		return math.Sqrt(float64(x)), nil
	} else {
		return 0, _sqrt_value(x)
	}
}
/***********/
//**********rot13密码
type _rot13_reader struct{
	r13 io.Reader
}
func _Rot13(b byte)byte{
	switch{
		case('A' <= b && b <= 'M')||('a' <= b && b<= 'm'):
		b += 13
		case('M' <= b && b <= 'Z')||('m' <= b && b<= 'z'):
		b -= 13
	}
	return b
}
func (r _rot13_reader)Read(b []byte)(int,error){
	n,err := r.r13.Read(b)
	for i := 0;i < n; i++{
		b[i] = _Rot13(b[i])
	}
	return n,err
}
/**********/
//**********图像
type Image struct{}
func (i Image)ColorModel() color.Model{
	return color.RGBAModel
}
func (i Image)Bounds()image.Rectangle{
	return image.Rect(0,0,200,200)
}
func (i Image)At(x,y int)color.Color{
	return color.RGBA{uint8(x),uint8(y),uint8(255),uint8(255)}
}
//**********

func main() {
	f1 := myfloat(-math.Sqrt(9))
	fmt.Println("f1 = ", f1.abs())

	vert1 := &vertex1{5, 5}
	fmt.Printf("转值之前 %+v ,平方和的开方 ：%v\n", vert1, vert1.abs())

	vert1.scale(10)
	fmt.Printf("转值之后 %+v, 平方和的开方：%v\n", vert1, vert1.abs())

	var ab abser
	f2 := myfloat(-math.Sqrt(9))
	vert2 := vertex1{3, 3}
	ab = f2
	ab = &vert2
	/*上一行为正确用法。下一句错误用法,因为以上abs用法只为指针变量定义
	ab = vert2
	*/
	fmt.Println("ab.abs() = ", ab.abs())

	var itf1 interf //接口interf
	var typ1 *typp
	itf1 = typ1
	describe(itf1)
	itf1.matt()

	itf1 = &typp{"hhh"}
	describe(itf1)
	itf1.matt()

	itf1 = F(math.Pi)
	describe(itf1)
	itf1.matt()

	var interf2 interface{} //空接口
	fmt.Printf("(%v, %T)\n", interf2, interf2)
	interf2 = 44
	fmt.Printf("(%v, %T)\n", interf2, interf2)
	interf2 = "gogogo"
	fmt.Printf("(%v, %T)\n", interf2, interf2)

	var interf3 interface{} = "helll"
	itf3 := interf3.(string)
	fmt.Println("itf3 = ", itf3)
	itf3, itf3_ok := interf3.(string)
	/***********
	   1.如果是float就赋值，2.如果是float值就返回true
	itf3,itf3_ok := interf3.(float64)
	***********/
	fmt.Println("itf3,itf3_ok = ", itf3, itf3_ok)

	_whattype(66)
	_whattype("this string")
	_whattype(false)

	person1 := person{"peter", 25}
	person2 := person{"jack", 31}
	fmt.Println("person1,person2 = ", person1, person2)

	host1 := map[string]_ipaddress{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range host1 {
		fmt.Printf("%v : %v \n", name, ip)
	}

	if error1 := _Run(); error1 != nil {
		fmt.Println(error1)
	}

	//输入负数、复数返回错误信息
	fmt.Println(_Sqrt(-1))
	fmt.Println(_Sqrt(4))

	my_reader1 := strings.NewReader("I am reader!,and this is a initual sentence")
	reader_bit := make([]byte, 8)
	for {
		n, err := my_reader1.Read(reader_bit)
		fmt.Printf("字节数：%d 错误信息：%v 字节内容：%v\n", n, err, reader_bit)
		fmt.Printf("reader_bit[:n] = %s \n", reader_bit[:n])
		if err == io.EOF {
			break
		}
	}

	my_reader2 := strings.NewReader("this is reader two!!")
	rot13_reader := _rot13_reader{my_reader2}
	io.Copy(os.Stdout,&rot13_reader)
	fmt.Println()

	img1 := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(img1.Bounds())
	fmt.Println(img1.At(0,0).RGBA())

	img2 := Image{}
	pic.ShowImage(img2)
}
