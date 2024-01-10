// 当表单不止一项时候
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(ctx *gin.Context) {
		var form LoginForm

		//使用了ShouldBindWith 而使用ShouldBind也可以
		if ctx.ShouldBindWith(&form, binding.Form) == nil {
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				ctx.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8092")
}
