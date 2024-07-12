package main

import (
	"fmt"
	"os"
)

// golang 移动文件
func main() {
	//err := os.Rename("C:\\Users\\zs\\AppData\\Roaming\\JetBrains\\IntelliJIdea2022.3.zip", "C:\\Users\\zs\\Desktop\\IntelliJIdea2022.3.zip")
	err := os.Rename("C:\\Users\\zs\\AppData\\Roaming\\JetBrains\\IntelliJIdea2022.3.zip", "C:\\网盘\\ali\\阿里云盘\\software\\idea配置以及插件\\IntelliJIdea2022.3.zip")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File moved successfully")
	}
}
