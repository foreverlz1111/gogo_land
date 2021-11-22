package main

import(
	"fmt"
	"math"
	"runtime"
	"time"
)
func sqrt(x float64)string {
     if x < 0{
     	return sqrt(-x)+"i"
     }
     return fmt.Sprint(math.Sqrt(x))
}
func pow(x,n,lim float64)float64{
    if v := math.Pow(x,n);v<lim{
       return v
       }else{
	fmt.Println("%g >= %g\n",v,lim)
	//v的寿命结束
	}
       return lim
}
func osrunning(){
     fmt.Print("运行的系统：")
     switch os := runtime.GOOS; os{
     case "darwin":
     fmt.Print("Mac OS \n ")
     case "linux":
     fmt.Print("Linux \n")
     case "freebsd":
     fmt.Print("openbsd \n")
     case "plan9":
     fmt.Print("Window\n")
     default:
     fmt.Print("%s \n",os)
     }
}
func whatsat(){
     fmt.Print("还有几天到周六")
     today := time.Now().Weekday()
     switch time.Saturday{
     case today + 0:
     fmt.Print(" 今天\n")
     case today + 1:
     fmt.Print(" 明天\n")
     case today + 2:
     fmt.Print(" 后天\n")
     default:
     fmt.Print("  遥遥无期\n")
     }
}
func sayhello(){
     t := time.Now()
     switch{
     case t.Hour() < 12:
     fmt.Println("早上好")
     case t.Hour() <17:
     fmt.Println("下午好")
     default:
     fmt.Println("晚上好")
     }
}
func main(){
     fmt.Println("hello world")
     sum := 0
     for i := 0;i<10;i++{
         //双分号三段式，条件可省略
     	 sum += i
	 }
     fmt.Println(sum)
     sum = 1
     for sum<1000{
         //等于while（sum<1000）
     sum += 1
     }
     fmt.Println(sum)
     /*下面里是无限循环
     for{
         fmt.Println()
     	}
     */
     fmt.Println(sqrt(2),sqrt(-3))     
     fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
		)
     osrunning()//当前运行的系统是
     whatsat()//什么时候星期六
     sayhello()//早上好中午好晚上好
     defer fmt.Println("world!")//推迟到外层函数返回执行
     fmt.Print("hello ")

     fmt.Println("运算中")
     for i := 0;i<10;i++{
     	 defer fmt.Println(i)
     }
     fmt.Println("运算完成")
     //defer 把输入投入栈中，先来的最后输出,可以做倒计时
}
