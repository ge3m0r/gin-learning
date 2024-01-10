package main

//当返回内容使用acsiijson 的时候会是什么样的情况

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言呀",
			"tag":  "<br/>",
		}
		ctx.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8080")
}

//经过测试，该方法存在问题，在谷歌，firefox上都会解析成相应的文字，所以显示跟json一样，
//但在safari浏览器上不会解析，edge 没有测试。
