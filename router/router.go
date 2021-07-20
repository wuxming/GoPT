package router

import (
	"gin1/controller"
	"gin1/middleware"
	"gin1/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//r.Use(middleware.CustomRouterMiddle1)
	r.Use(utils.CORS(utils.Options{Origin: "*"})) //跨域
	Test := r.Group("/test")
	Test.Use(middleware.CustomRouterMiddle1) //中间件按顺序走的
	{
		Test.GET("/find", controller.Find, middleware.Verification)
		Test.POST("/create", controller.Create)
		Test.PUT("/update", controller.Update)
		Test.DELETE("/delete", controller.Delete)
		Test.GET("/findById", controller.FindById)
	}
	return r
}
