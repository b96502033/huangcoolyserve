package dao

import (
	"context"
	"fmt"
	"huangcoolyserver/src/model"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ComUserNavbarDao struct {
	MgoBaseDao
	collection *mongo.Collection
	onceC      sync.Once
}

func NewComUserNavbarDao() *ComUserNavbarDao {
	comUserNavbarDao := new(ComUserNavbarDao)
	comUserNavbarDao.onceC.Do(func() {
		collection := getCollection("ComUserNavbar")
		comUserNavbarDao.collection = collection
	})
	return comUserNavbarDao
}

func (e *ComUserNavbarDao) GetComUserNavbarById(userId primitive.ObjectID) ([]*model.ComUserNavbar, error) {
	fmt.Println(userId)
	var comUserNavbarSlice []*model.ComUserNavbar
	filter := bson.D{primitive.E{Key: "ComUserId", Value: userId}}
	multiResult, _ := e.collection.Find(context.TODO(), filter, options.Find())
	if multiResult.Err() != nil {
		return nil, multiResult.Err()
	}
	for multiResult.Next(context.TODO()) {
		//定義一個文件，將單個文件解碼為result
		var result model.ComUserNavbar
		multiResult.Decode(&result)
		comUserNavbarSlice = append(comUserNavbarSlice, &result)
	}

	return comUserNavbarSlice, nil
}
