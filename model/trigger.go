package model

import "github.com/globalsign/mgo/bson"

type Trigger struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name" bw:"name"`
	Enable      bool          `json:"enable" bson:"enable"`
	NextTime    uint64        `json:"next_time" bson:"next_time"`
	Parallel    int           `json:"parallel" bson:"parallel"`
	TriggerType int           `json:"trigger_type" bson:"trigger_type"`
	TypeID      bson.ObjectId `json:"type_id" bson:"type_id"`
	CreateTime  string        `json:"create_time" bson:"create_time"`
}

//TriggerType
//The notice way to consumer
//Kind:
// 0 - http.
// 1 - nsq.
type TriggerType struct {
	ID       bson.ObjectId     `json:"_id" bson:"_id"`
	Kind     int               `json:"kind" bson:"kind"`
	Endpoint string            `json:"endpoint" bson:"endpoint"`
	Ext      map[string]string `json:"ext" bson:"ext"`
}
