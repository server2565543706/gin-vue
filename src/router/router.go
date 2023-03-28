package router

import (
	"github.com/gin-gonic/gin"
	"github.com/server2565543706/gin-vue/src"
)

func Run() {
	r := gin.Default()

	r.GET("/api/user/register", src.Register)

	r.Run()
}
