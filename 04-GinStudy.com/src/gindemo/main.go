package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()
	// 配置路由
	engine.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hell Gin"})
	})

	engine.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "post request"})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	engine.Run(":8000")
}
