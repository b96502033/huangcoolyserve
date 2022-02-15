package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComUser struct {
	Id           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserName     string             `json:"userName,omitempty"`
	UserWrorkId  string             `json:"userWrorkId,omitempty"`
	UserPassword string             `json:"userPassword,omitempty"`
	Login        bool               `json:"login,omitempty"`
	LoginKey     string             `json:"loginKey,omitempty"`
}
