package controller

import (
	"gin1/model"
	"gin1/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func Find(c *gin.Context) {
	log.Println("find===================")
	result := service.Find()
	c.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
	var book model.Test
	c.Bind(&book)
	result := service.Create(book)
	c.JSON(http.StatusOK, result)
}
func Update(c *gin.Context) {
	var book model.Test
	c.Bind(&book) //这个会取到postman中json的值
	result := service.Update(book)
	c.JSON(http.StatusOK, result)
}
func Delete(c *gin.Context) {
	_id, _ := primitive.ObjectIDFromHex(c.Query("_id")) //转换为objectid
	result := service.Delete(_id)
	c.JSON(http.StatusOK, result)
}
func FindById(c *gin.Context) {
	_id, _ := primitive.ObjectIDFromHex(c.Query("_id"))
	result := service.FindById(_id)
	c.JSON(http.StatusOK, result)
}
