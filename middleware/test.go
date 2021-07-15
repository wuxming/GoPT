package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func CustomRouterMiddle1(c *gin.Context)  {
	t := time.Now()
	log.Println("我是自定义中间件第1种定义方式---请求之前")
	c.Set("example","CustomRouterMiddle1")
	c.Next()
	log.Println("我是自定义中间件第1种定义方式---请求之后")
	t2 := time.Since(t)
	log.Println(t2)
}
func Verification(c *gin.Context) {
	url := c.Request.URL//获取url  /test/find
	log.Println(url,"-------------------")
	method := c.Request.Method//获取 方法
	log.Println(method,"----------------")
	header := c.Request.Header.Get("token") //获取请求头上的数据
	log.Println(header,"header")
	c.Abort()
}
