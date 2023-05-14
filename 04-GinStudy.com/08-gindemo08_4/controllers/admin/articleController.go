package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	// c.String(http.StatusOK, "文章列表--")
	con.success(c)
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "-Edit---文章---")
}
