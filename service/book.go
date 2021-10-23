package service

import (
	"context"
	"gin1/model"
	"gin1/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = utils.Coll().Collection("libraries") //连接到哪个表
func Find() utils.Response {
	cur, err := collection.Find(context.TODO(), bson.M{}) //查询所有 M map没有顺序 Binary JSON
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	var books []*model.Test
	err = cur.All(context.TODO(), &books)
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	return utils.SuccessMess("All book", books)
}
func Create(s1 model.Test) utils.Response {
	createResult, err := collection.InsertOne(context.Background(), s1)
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	return utils.SuccessMess("创建成功", createResult)
}
func Update(s1 model.Test) utils.Response { //根据id更新
	filter := bson.M{"_id": s1.Id} //查找的对象
	update := bson.M{"$set": s1}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	return utils.SuccessMess("更新成功", updateResult)
}
func Delete(_id primitive.ObjectID) utils.Response {
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": _id})
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	return utils.SuccessMess("删除成功", deleteResult)
}
func FindById(_id primitive.ObjectID) utils.Response {
	var book model.Test
	err := collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&book)
	if err != nil {
		return utils.ErrorMess("查找失败", nil)
	}
	return utils.SuccessMess("查找成功", book)
}
