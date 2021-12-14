package main

import(
	"fmt"
	"math"
	
)
type vertex1 struct{
	X, Y float64
}
type myfloat float64
func (v vertex1)abs()float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func abs(v vertex1)float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func (myf myfloat)abs()float64{
	if myf < 0{
		return float64(-myf)
	}
	return float64(myf)
}
func main(){
	v1 := vertex1{5,7}
	fmt.Println("vertex1(5,7) = ",v1.abs())

	v2 := vertex1{3,4}
	fmt.Println("vertex1(3,4) = ",abs(v2))

	f1 := myfloat(-math.Sqrt2)
	fmt.Println("f1 = ",f1.abs())
}
