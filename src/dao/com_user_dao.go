package dao

import (
	"context"
	"huangcoolyserver/src/model"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ComUserDao struct {
	MgoBaseDao
	collection *mongo.Collection
	onceC      sync.Once
}

func NewComUserDao() *ComUserDao {
	comUserDao := new(ComUserDao)
	comUserDao.onceC.Do(func() {
		collection := getCollection("ComUser")
		comUserDao.collection = collection
	})

	return comUserDao
}

func (e *ComUserDao) GetComUserByUserWorkId(userWorkId string) (*model.ComUser, error) {

	var comUser model.ComUser
	filter := bson.D{primitive.E{Key: "UserWorkId", Value: userWorkId}}
	singleResult := e.collection.FindOne(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	err := singleResult.Decode(&comUser)
	return &comUser, err
}

func (e *ComUserDao) GetComUserByUserId(userId primitive.ObjectID) (*model.ComUser, error) {

	var comUser model.ComUser
	filter := bson.D{primitive.E{Key: "_id", Value: userId}}
	singleResult := e.collection.FindOne(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	err := singleResult.Decode(&comUser)
	return &comUser, err
}

func (e *ComUserDao) UpdateComUserLogin(userId primitive.ObjectID, login bool) error {

	setLogin := primitive.E{Key: "login", Value: login}
	setE := primitive.E{Key: "$set", Value: bson.D{setLogin}}

	result, err := e.collection.UpdateMany(
		context.TODO(),
		bson.M{"_id": userId},
		bson.D{
			setE,
		},
	)

	if err == nil {
		log.Println("Updated %i Documents!\n", result.ModifiedCount)
	}
	return err

}
