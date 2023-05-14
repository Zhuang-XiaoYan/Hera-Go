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

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}

	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-userlist")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-plist")
		})
	}

	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "新闻列表")
		})
	}

	r.Run()
}
