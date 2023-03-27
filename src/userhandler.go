package src

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
)

func Register(ctx *gin.Context) {
	// 获取参数
	// Get the parameter
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	// data verification

	// Create user

	// Return result

	ctx.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success!",
	}))
}
