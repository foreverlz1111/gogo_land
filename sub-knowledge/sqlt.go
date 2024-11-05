package main

import (
	"context"
	"log"
	"time"
)

func main() {
	var num int
	num = 913936
	left, right := 0, num
	ctx, cancelctx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelctx()
	for {
		select {
		case <-ctx.Done(): // 检查是否超时
			log.Println("操作超时")
			return
		default:
			mid := (left + right) / 2
			if mid*mid == num { // 找到结果
				log.Println("平方根为:", mid)
				return
			} else if mid*mid > num {
				right = mid
			} else {
				left = mid
			}
			if right-left <= 1 {
				log.Println("未找到精确的平方根，近似值为:", left)
				return
			}
		}
	}
}
