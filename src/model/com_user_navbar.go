package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComUserNavbar struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ComUserId   primitive.ObjectID `json:"comUserId,omitempty"`
	SysNavbarId primitive.ObjectID `json:"sysNavbarId,omitempty"`
}
