package main

import (
	"fmt"
	"gindemo06/routers"
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

	// 应用其他的文件routers
	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
