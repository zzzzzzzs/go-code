// 循环 100次

package main

import (
	"fmt"
	"time"
)

func main() {
	// 记录时间
	start := time.Now()
	// 数字
	var num = 0
	for i := 0; i < 10000000000; i++ {
		num++
	}
	println(num)
	// 计算时间
	end := time.Now()
	fmt.Println(end.Sub(start))
}
