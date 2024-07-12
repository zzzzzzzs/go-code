package main

import "github.com/gin-gonic/gin"

// 这是一个 gin 测试代码
func main() {
	// 创建一个服务
	ginServer := gin.Default()

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello"})
	})

	ginServer.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "这是一个测试程序"})
	})

	ginServer.Run(":8082")
}
