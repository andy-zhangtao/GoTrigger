package model

import "github.com/globalsign/mgo/bson"

type Trigger struct {
	ID         bson.ObjectId `json:"_id" bson:"_id" bw:"_id"`
	Name       string        `json:"name" bson:"name" bw:"name"`
	Async      bool          `json:"async" bson:"async"`
	Enable     bool          `json:"enable" bson:"enable"`
	Interval   int           `json:"interval" bson:"interval"`
	NextTime   int64         `json:"next_time" bson:"next_time"`
	Parallel   int           `json:"parallel" bson:"parallel"`
	Type       TriggerType   `json:"type" bson:"type"`
	CreateTime string        `json:"create_time" bson:"create_time"`
	UpdateTime string        `json:"update_time" bson:"update_time"`
}

//TriggerType
//The trigger plugin
//Kind:
// 0 - http.
// 1 - nsq.
// ...
type TriggerType struct {
	Kind     int               `json:"kind" bson:"kind"`
	Endpoint string            `json:"endpoint" bson:"endpoint"`
	Ext      map[string]string `json:"ext" bson:"ext"`
}

type TriggerPlugin struct {
	ID   bson.ObjectId `json:"_id" bson:"_id" bw:"_id"`
	Name string        `json:"name" bson:"name" bw:"name"`
	PID  int           `json:"pid" bson:"pid" bw:"pid"`
	Desc string        `json:"desc" bson:"desc"`
}
