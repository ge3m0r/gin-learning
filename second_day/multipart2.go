package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "anonymous")
		ctx.JSON(200, gin.H{
			"stautus":  "posted",
			"message":  message,
			"nickname": nick,
		})
	})
	router.Run(":8081")
}
