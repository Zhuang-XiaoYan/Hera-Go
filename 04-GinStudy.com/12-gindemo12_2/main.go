package main

import (
	"gindemo12/models"
	"gindemo12/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")
	// 创建基于cookie的存储引擎。secrt111是用于加密的密钥
	store := cookie.NewStore([]byte("secrt111"))
	// 配置session的中间件store是前面船舰的存储引擎。我们也九二一替换成为其他的引擎。
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run(":80")
}
