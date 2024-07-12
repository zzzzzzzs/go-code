package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO 此程序是离线分析日志中包含 error 的程序，并能将 offset 保存下来，下次程序启动的时候从 offset 的位置分析就可以了。

const offsetFile = "offset.txt"

func main() {
	offset := readOffset()

	file, err := os.Open("astaxie.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.Seek(offset, 0); err != nil {
		fmt.Println("Error seeking to offset:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		offset += int64(len(line)) + 1
		if strings.Contains(line, "error") {
			fmt.Println("Error:", line)
		}
		//fmt.Println(line)
		writeOffset(offset)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning log file:", err)
		return
	}
}

func readOffset() int64 {
	file, err := os.Open(offsetFile)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	offset, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return 0
	}
	return offset
}

func writeOffset(offset int64) {
	file, err := os.Create(offsetFile)
	if err != nil {
		fmt.Println("Error creating offset file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(strconv.FormatInt(offset, 10)); err != nil {
		fmt.Println("Error writing offset:", err)
		return
	}
}
