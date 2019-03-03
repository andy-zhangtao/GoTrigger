package util

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
)

var triggerChan chan bson.ObjectId
var fireChan chan model.Trigger

func InitTriggerChan() {
	triggerChan = make(chan bson.ObjectId)
	fireChan = make(chan model.Trigger)
}

func GetTriggerChan() chan bson.ObjectId {
	return triggerChan
}

func GetFireChan() chan model.Trigger {
	return fireChan
}
