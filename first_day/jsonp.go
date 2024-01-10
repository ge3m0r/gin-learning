// 向不同数据域中请求数据，如果参数存在回调，将回调添加到响应体中
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		//如果参数不是callback那么将返回原来的值
		c.JSONP(http.StatusOK, data)
	})
	r.Run(":8080")
}
