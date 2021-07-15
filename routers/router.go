package routers

import (
	"gin1/controller"
	"gin1/cors"
	"gin1/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.New()
	//r.Use(middleware.CustomRouterMiddle1)
	r.Use(cors.CORS(cors.Options{Origin: "*"}))//跨域
	Test := r.Group("/test")
	Test.Use(middleware.CustomRouterMiddle1)//中间件按顺序走的
	{
		Test.GET("/find",middleware.Verification,controller.Find)
		Test.POST("/create",controller.Create)
		Test.PUT("/update",controller.Update)
		Test.DELETE("/delete",controller.Delete)
		Test.GET("/findById",controller.FindById)
	}
	return  r
}