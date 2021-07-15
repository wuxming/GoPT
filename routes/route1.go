package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "666",
	})
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}
func main() {
	//1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello.go")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello www.wxm.com!",
		})
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")     //使用这个
		action := c.Param("action") //这个打印的时候有前面的斜杠
		c.String(http.StatusOK, name)
		c.String(http.StatusOK, action)
	})
	r.GET("/index", func(c *gin.Context) {
		name := c.DefaultQuery("name", "") //设定默认参数
		c.String(http.StatusOK, name)
		c.String(http.StatusOK, fmt.Sprint(name))
	})
	//表单
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, types+"\n"+username+" "+password)
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
