package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` //omitempty 如果为空就没有这个字段
	Bookname    string `bson:"bookname" json:"bookname"` //这里要注意是空格不是，
	Author      string `bson:"author" json:"author"`//结构体标签
	Instruction string `bson:"instruction" json:"instruction"`
	Img         string `bson:"img" json:"img"`
}