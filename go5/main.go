package main
import(
	"fmt"
)
type vertex struct{
     X int
     Y int
     }
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
     fmt.Println("v.X = ",vert.X)
}