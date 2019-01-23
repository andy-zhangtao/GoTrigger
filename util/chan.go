package util

import "github.com/globalsign/mgo/bson"

var triggerChan chan bson.ObjectId

func InitTriggerChan() {
	triggerChan = make(chan bson.ObjectId)
}

func GetTriggerChan() chan bson.ObjectId {
	return triggerChan
}
