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
}