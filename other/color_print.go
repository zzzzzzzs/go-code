package main

import "fmt"

func main() {
	// 带颜色打印
	fmt.Println("\033[1;32mHello World\033[0m")
	// 带红色打印
	fmt.Println("\033[1;31mHello World\033[0m")
}
