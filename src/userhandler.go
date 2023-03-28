package src

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
	"net/http"
)

func Register(ctx *gin.Context) {
	// 获取参数
	// Get the parameter
	//name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	// data verification
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The Phone num must be 11 digits!",
		})
	}
	//verification password
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "password cannot be less than 6 digits",
		})
	}
	// Create user

	// Return result

	ctx.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success!",
	}))
}
