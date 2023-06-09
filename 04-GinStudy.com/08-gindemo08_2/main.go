package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 多个中间件的执行流程是一个嵌套的过程
func initMiddlewareOne(c *gin.Context) {

	fmt.Println("1-我是一个中间件-initMiddlewareOne")
	// 停止调用的路由的中的请求程序
	c.Next()
	fmt.Println("2-我是一个中间件-initMiddlewareOne")

}
func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1-我是一个中间件-initMiddlewareTwo")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("2-我是一个中间件-initMiddlewareTwo")

}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
		fmt.Println("这是一个首页")
		c.String(200, "gin首页")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "新闻页面")
	})
	r.GET("/login", func(c *gin.Context) {
		c.String(200, "login")
	})

	r.Run()
}
