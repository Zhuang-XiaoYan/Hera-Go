package routers

import (
	"gindemo07/controllers/index"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/news", index.DefaultController{}.News)
	}
}
