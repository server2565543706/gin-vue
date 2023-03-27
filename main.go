package main

import (
	"github.com/gin-gonic/gin"
	"github.com/server2565543706/gin-vue/src"
)

func main() {
	r := gin.Default()
	r.GET("/ping", src.Ping)
	r.GET("/api/user/register", src.Register)

	r.Run()
}
