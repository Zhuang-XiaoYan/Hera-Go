package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	// 设置session
	session := sessions.Default(c)
	session.Set("username", "张三")
	session.Save()
	con.success(c)
}
func (con ArticleController) Add(c *gin.Context) {
	// 获取session
	session := sessions.Default(c)
	userName := session.Get("username")
	fmt.Printf("session value=%v", userName)
	c.String(http.StatusOK, "-add--文章-")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "-Edit---文章---")
}
