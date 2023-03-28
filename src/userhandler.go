package src

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	// 获取参数
	// Get the parameter
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if len(phone) != 11 { // data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The Phone num must be 11 digits!",
		})
		return
	}

	if len(password) < 6 { //verification password
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "password cannot be less than 6 digits",
		})
		return
	}

	if len(name) == 0 { // If npt input the name, we can give  a 10-digit string
		name = RandomString(10)

	}

	log.Println(name, password, phone)
	// Create user

	ctx.JSON(200, gin.H{ // Return result
		"msg": "Register success!",
	})
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k, _ := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
