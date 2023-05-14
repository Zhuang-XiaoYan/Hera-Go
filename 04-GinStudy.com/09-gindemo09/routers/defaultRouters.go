package routers

import (
	"gindemo09/controllers/index"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/news", index.DefaultController{}.News)
	}
}
