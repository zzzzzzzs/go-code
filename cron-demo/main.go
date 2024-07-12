package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	// 支持秒
	c := cron.New(cron.WithSeconds())
	EntryID, err := c.AddFunc("* * * * * *", func() {
		fmt.Println("asdasdasd")
	})
	if err != nil {
		fmt.Println(EntryID, err)
	}
	c.Start()

	// 程序将在后台运行，不会退出
	select {}
}
