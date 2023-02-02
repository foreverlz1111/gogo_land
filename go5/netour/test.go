package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	has := crc32.NewIEEE()
	wri := io.MultiWriter(has, os.Stdout)
	fmt.Fprintf(wri, "hello world\n")
	fmt.Printf("hash = %#x\n", has.Sum32())
	fmt.Println(time.Since(start))

	respond, err := http.Get("https://www.163.com")
	if err != nil {
		fmt.Printf("%s\n", err)
		fmt.Println(time.Since(start))
		return
	}
	defer respond.Body.Close()
	respondbody,err := io.ReadAll(respond.Body)
	fmt.Println(time.Since(start))
	fmt.Printf("%s",respondbody)
	//	client := &http.Client{
	//		CheckRedirect: redirectPolicyFunc,
	//	}

}
