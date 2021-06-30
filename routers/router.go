package routers

import (
	"gin1/controller"
	"gin1/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.CORS(cors.Options{Origin: "*"}))//跨域
	Test := r.Group("/test")
	{
		Test.GET("/find",controller.Find)
		Test.POST("/create",controller.Create)
		Test.PUT("/update",controller.Update)
		Test.DELETE("/delete",controller.Delete)
	}
	return  r
}