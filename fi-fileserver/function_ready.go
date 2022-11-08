package main

import (
	"fmt"
	"os"
	"strconv"
)

func bmgt(int642 int64) string {
	switch {
	default:
		return ""
	case int642 < int64(1024):
		return strconv.FormatInt(int64(int642), 10) + "B"
	case int642 < int64(1024*1024):
		return strconv.FormatInt(int64(int642/1024), 10) + "KB"
	case int642 < int64(1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024), 10) + "MB"
	case int642 < int64(1024*1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024/1024), 10) + "GB"
	}

}
func main() {

	fmt.Println("test:")
	cur, _ := os.Getwd()
	entry, _ := os.ReadDir(cur)
	for _, e := range entry {
		t, _ := e.Info()
		fmt.Println(e.Name(), bmgt(t.Size()), t.ModTime().Format("2006年01月02日 15:04:05"))

		//os.Stat(e.Name())
	}
}
