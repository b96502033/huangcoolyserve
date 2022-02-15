package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SysNavbar struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	NavbarName string             `json:"navbarName,omitempty"`
}
