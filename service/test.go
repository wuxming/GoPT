package service

import (
	"context"
	"fmt"
	"gin1/config"
	"gin1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

var collection =config.Coll().Collection("libraries")//连接到哪个表
func Find()([]*models.Test,int)  {
	cur, err :=collection.Find(context.Background(),bson.M{})//查询所有 M map没有顺序 Binary JSON
	if err != nil {
		log.Fatal(err)
	}
	if err:=cur.Err();err!=nil {
		log.Fatal(err)
	}
	var all  []*models.Test
	err = cur.All(context.Background(),&all)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	code := http.StatusOK
	fmt.Println(all)
	return all,code
}
func Create(s1 models.Test) int  {
	var code int = http.StatusOK
	createResult,err :=collection.InsertOne(context.Background(),s1)
	fmt.Println(createResult)
	if err != nil {
		log.Fatal(err)//出错程序停止
		code = http.StatusInternalServerError
	}
	return  code
}
func Update(s1 models.Test)int{
	fmt.Println(s1.Id)
	filter := bson.M{"_id":s1.Id} //查找的对象
	update := bson.M{"$set":s1}
	updateResult,err :=collection.UpdateOne(context.Background(),filter,update)
	fmt.Println(updateResult)
	if err != nil {
		log.Fatal(err)
	}
	return http.StatusOK
}
func Delete(id primitive.ObjectID)int{
	code :=http.StatusOK
	fmt.Println(id)
	deleteResult,err :=collection.DeleteOne(context.Background(),bson.M{"_id":id})
	if err!=nil {
		code = http.StatusInternalServerError
		log.Fatal(err)
	}
	fmt.Println(deleteResult)
	return code
}