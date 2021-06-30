package controller

import (
	"gin1/models"
	"gin1/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Find(c *gin.Context)  {
	b,_ := service.Find()
	c.Bind(&b)
	c.JSON(http.StatusOK,gin.H{
		"data":b,
	})
}

func Create(c *gin.Context)  {
	var book models.Test
	c.Bind(&book)
	code := service.Create(book)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
	})
}
func Update(c *gin.Context)  {
	var book models.Test
	c.Bind(&book)//这个会取到postman中json的值
	code := service.Update(book)
	c.JSON(code,gin.H{
		"code":code,
	})
}
func Delete(c *gin.Context)  {
	id,_ :=primitive.ObjectIDFromHex(c.Query("_id"))
	code := service.Delete(id)
	c.JSON(code,gin.H{
		"code":code,
	})
}