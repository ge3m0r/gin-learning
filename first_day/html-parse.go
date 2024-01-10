package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	/*router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")*/

	/*router := gin.Default()
	//经过测试templates/ 中间使用一颗星还是两颗星都可以运行
	router.LoadHTMLGlob("templates/**/ /*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8088")*/

	//可以自定义分隔符和自定义模版渲染
	router := gin.Default()
	//自定义分隔符
	router.Delims("{[{", "}]}")
	//设置模版函数
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//加载模版文件
	router.LoadHTMLFiles("./templates/raw.tmpl")
	//模版写法上的问题：.now 是运算的一个结果，如果将这个结果传入另一个参数可以使用 ｜ 传入也可以
	//函数名 参数直接运行，如果这个模版参数可控会有安全问题
	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8090")

}
