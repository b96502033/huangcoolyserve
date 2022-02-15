package dao

import (
	"context"
	"huangcoolyserver/src/model"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SysNavbarDao struct {
	MgoBaseDao
	collection *mongo.Collection
	onceC      sync.Once
}

func NewSysNavbarDao() *SysNavbarDao {
	sysNavbarDao := new(SysNavbarDao)
	sysNavbarDao.onceC.Do(func() {
		collection := getCollection("SysNavbar")
		sysNavbarDao.collection = collection
	})

	return sysNavbarDao
}

func (e *SysNavbarDao) GetSysNavbarById(sysNavbarId primitive.ObjectID) (*model.SysNavbar, error) {

	var sysNavbar model.SysNavbar
	filter := bson.D{primitive.E{Key: "_id", Value: sysNavbarId}}
	singleResult := e.collection.FindOne(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	err := singleResult.Decode(&sysNavbar)
	return &sysNavbar, err
}
