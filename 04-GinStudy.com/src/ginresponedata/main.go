package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 配置路由
	engine.GET("/test1", func(c *gin.Context) {
		c.String(http.StatusOK, "返回值：%v", "test")
	})

	engine.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "post request"})
	})

	engine.GET("/test3", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "yes",
			"code":   200,
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	engine.Run(":8080")
}
