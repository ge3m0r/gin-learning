package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化。对于gin 进行初始化
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		//返回的内容是 json 格式，同时是在 http status 是 200 的时候返回
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
