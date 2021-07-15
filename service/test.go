package service

import (
	"context"
	"gin1/config"
	"gin1/helper"
	"gin1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection =config.Coll().Collection("libraries")//连接到哪个表
func Find()map[string]interface{}  {
	cur, err :=collection.Find(context.TODO(),bson.M{})//查询所有 M map没有顺序 Binary JSON
	if err != nil {
		return helper.ErrorMess(err.Error(),nil)
	}
	var books  []*models.Test
	err = cur.All(context.TODO(),&books)
	if err != nil {
		return helper.ErrorMess(err.Error(),nil)
	}
	return helper.SuccessMess("All book",books)
}
func Create(s1 models.Test) map[string]interface{}  {
	createResult,err :=collection.InsertOne(context.Background(),s1)
	if err != nil {
		return helper.ErrorMess(err.Error(),nil)
	}
	return   helper.SuccessMess("创建成功",createResult)
}
func Update(s1 models.Test)map[string]interface{}{//根据id更新
	filter := bson.M{"_id":s1.Id} //查找的对象
	update := bson.M{"$set":s1}
	updateResult,err :=collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		return helper.ErrorMess(err.Error(),nil)
	}
	return helper.SuccessMess("更新成功",updateResult)
}
func Delete(_id primitive.ObjectID)map[string]interface{}{
	deleteResult,err :=collection.DeleteOne(context.Background(),bson.M{"_id":_id})
	if err!=nil {
		return helper.ErrorMess(err.Error(),nil)
	}
	 return helper.SuccessMess("删除成功",deleteResult)
}
func FindById(_id primitive.ObjectID) map[string]interface{} {
	var  book models.Test
	err := collection.FindOne(context.TODO(),bson.M{"_id":_id}).Decode(&book)
	if err != nil {
		return helper.ErrorMess("查找失败",nil)
	}
	return  helper.SuccessMess("查找成功",book)
}