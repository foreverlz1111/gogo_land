package main

import(
	"fmt"
	"time"
	"hash/crc32"
	"os"
	"io"
	"net/http"
)

func main(){
	start := time.Now()
	has := crc32.NewIEEE()
	wri := io.MultiWriter(has,os.Stdout)
	fmt.Fprintf(wri,"hello world\n")
	fmt.Printf("hash = %#x\n",has.Sum32())
	fmt.Println(time.Since(start))

	read,err := http.Get("https://www.163.com")
	if err != nil{
		fmt.Printf("%s\n",err)
		fmt.Println(time.Since(start))
		return
	}
	fmt.Println(read)
	fmt.Println(time.Since(start))
	
}
