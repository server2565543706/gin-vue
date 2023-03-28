package router

import (
	"github.com/gin-gonic/gin"
	"github.com/server2565543706/gin-vue/src"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")

	api := r.Group("v1")
	{
		api.POST("/user/register", src.Register)

	}

	return r
}
