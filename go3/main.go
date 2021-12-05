package main
import(
	"fmt"
	"math/rand"
	"time"
)
type vertex struct{
     X int
     Y int
     }
var(
	v1 = vertex{2,3}
	v2 = vertex{X:1}
	v3 = vertex{}
	v4 = &vertex{1,2}
     )
var array1 [8]int
var str1 [2]string

func main(){
     fmt.Println("hello pointer")
     i,j := 53,443
     p := &i
     fmt.Println("*p = ",*p)
     *p = 21
     fmt.Println("i = ",i)

     p = &j
     *p = *p / 443
     fmt.Println("j = ",j)

     fmt.Println("struct :",vertex{3,7})

     vert := vertex{3,7}
     vert.X = 5
     vv := &vert
     vv.X = 1e9
     fmt.Println("v.X = ",vert.X)
     fmt.Println("vv.x= ",vv.X)

     fmt.Println(v1,v2,v3,v4)

     str1[0] = "Hello"
     str1[1] = "GOlang"
     for a := 0;a<len(str1);a++{
     	 fmt.Println("str1[",a,"]= ",str1[a])
	 }
	 fmt.Println("array1 int len() = ",len(array1))
	 r := rand.New(rand.NewSource(time.Now().UnixNano()))
	 for a := 0;a<len(array1);a++{
	 array1[a] = r.Intn(66)
	 }
	 array2 := array1[4:7]//array2[4]、array2[5]、array2[6]这三个
	 fmt.Println("array1 = ",array1)
	 fmt.Println("array2 = ",array2)

      names := [4]string{"name1","name2","name3","name4"}//[4]中的数字4可以省略
      fmt.Println(names)
      name_b := names[1:3]
      fmt.Println(names[0:2],name_b)

      array3 := []int{1,2,3,4,5,6,7,8}
      fmt.Println("array3 = ",array3[0:3])
      array4 := []bool{true,true,false}
      fmt.Println("array4",array4)
      array5 := []struct{
      a int
      b bool}{
      {1,true},
      {2,false},
      {3,true},
      }//严格的缩进
      fmt.Println("array5",array5)
      
      printSlice(array3)
      
      s := array3[:0]//取0个
      printSlice(s)
      s = s[:4]//取四个
      printSlice(s)
      s = s[2:]//舍弃前两个
      printSlice(s)

      //s = s[6:6]//此时s也是空len()和空cap()但不显示为nil
      
      var array_emp []int
      fmt.Println(array_emp, len(array_emp), cap(array_emp))
      if  array_emp == nil {
      	 fmt.Println("空数组！")
	 }
}
func printSlice(s []int){
     fmt.Printf("长度= %d ，首个元素 = %d s= %v\n",len(s),cap(s),s)
}