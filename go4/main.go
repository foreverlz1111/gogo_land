package main

import(
	"fmt"
	"math"
	
)
type vertex1 struct{
	X, Y float64
}
type myfloat float64
//****************************************
func (v *vertex1)abs()float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
/* 非指针：
   func abs(v vertex1)float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
        }
*****************************************/
func (v *vertex1)scale(f float64){
	v.X = v.Y * f
	v.Y = v.X * f
}
//****************************************
func scale(v *vertex1,f float64){
	v.X = v.Y * f
	v.Y = v.X * f
}
/* 非指针：
   func scale(v vertex1,f float64){
	v.X = v.Y * f
	v.Y = v.X * f
        }
*****************************************/
func (myf myfloat)abs()float64{
	if myf < 0{
		return float64(-myf)
	}
	return float64(myf)
}
type abser interface{
	abs() float64
}
//********************1接口——2结构——3方法
type interf interface{
	matt()
}
type typp struct{
	str string
}
func (t *typp) matt(){
	if t == nil{
		fmt.Println("t的值为 null!")
		return
	}
	fmt.Println(t.str)
	/**********非指针
	  func (t typp) matt(){
	  fmt.Println(t.str)
	  }
	  **********/
}/********************/
func (f F) matt(){
	fmt.Println("f = ",f)
}
type F float64
func describe(i interf) {
	fmt.Printf("(数值：%v,类型： %T)\n", i, i)
}
func _whattype(i interface{}){
	switch v := i.(type){
		case int:fmt.Printf(" %v ,乘以二 %v\n",v,v*2)
		case string:fmt.Printf("%q 类型是 , 共 %v 字节\n",v,len(v))
		default:fmt.Printf("未定义识别到的类型，(%T) \n",v)
	}
}
type person struct {
	_name string
	_age int
}
func (p person) String() string{
	return fmt.Sprintf("%v (%v years)",p._name,p._age)
}
func main(){
	f1 := myfloat(-math.Sqrt(9))
	fmt.Println("f1 = ",f1.abs())

	vert1 := &vertex1{5,5}
	fmt.Printf("转值之前 %+v ,平方和的开方 ：%v\n",vert1,vert1.abs())
	
	vert1.scale(10)
	fmt.Printf("转值之后 %+v, 平方和的开方：%v\n",vert1,vert1.abs())

	var ab abser
	f2 := myfloat(-math.Sqrt(9))
	vert2 := vertex1{3,3}
	ab = f2
	ab = &vert2
	/*上一行为正确用法。下一句错误用法,因为以上abs用法只为指针变量定义
	ab = vert2
	*/
	fmt.Println("ab.abs() = ",ab.abs())

	var itf1 interf//接口interf
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

	var interf2 interface{}//空接口
	fmt.Printf("(%v, %T)\n", interf2, interf2)
	interf2 = 44
	fmt.Printf("(%v, %T)\n", interf2, interf2)
	interf2 = "gogogo"
	fmt.Printf("(%v, %T)\n", interf2, interf2)

	var interf3 interface{} = "helll"
	itf3 := interf3.(string)
	fmt.Println("itf3 = ",itf3)
	itf3,itf3_ok := interf3.(string)
	/***********
	   1.如果是float就赋值，2.如果是float值就返回true
	itf3,itf3_ok := interf3.(float64)
	***********/
	fmt.Println("itf3,itf3_ok = ",itf3,itf3_ok)

	_whattype(66)
	_whattype("this string")
	_whattype(false)

	person1 := person{"peter",25}
	person2 := person{"jack",31}
	fmt.Println("person1,person2 = ",person1,person2)

}
