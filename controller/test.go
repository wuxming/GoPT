package controller

import (
	"gin1/models"
	"gin1/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func Find(c *gin.Context)  {
	log.Println("find===================")
	result := service.Find()
	c.JSON(result["status"].(int),result)
}

func Create(c *gin.Context)  {
	var book models.Test
	c.Bind(&book)
	result := service.Create(book)
	c.JSON(result["status"].(int),result)
}
func Update(c *gin.Context)  {
	var book models.Test
	c.Bind(&book)//这个会取到postman中json的值
	result := service.Update(book)
	c.JSON(result["status"].(int),result)
}
func Delete(c *gin.Context)  {
	_id,_ :=primitive.ObjectIDFromHex(c.Query("_id"))//转换为objectid
	result := service.Delete(_id)
	c.JSON(result["status"].(int),result)
}
func FindById(c *gin.Context)  {
	_id,_ :=primitive.ObjectIDFromHex(c.Query("_id"))
	result := service.FindById(_id)
	c.JSON(result["status"].(int),result)
}