package model

import "github.com/globalsign/mgo/bson"

type TriggerJnl struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Time    string        `json:"time" bson:"time"`
	Status  int           `json:"status" bson:"status"`
	Message string        `json:"message" bson:"message"`
}
