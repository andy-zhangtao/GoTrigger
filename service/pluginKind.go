package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
)

func AddNewPluginKind(plugin model.PluginType) (model.PluginType, error) {
	plugin.ID = bson.NewObjectId()
	plugin.PID = db.GetNextSeq()
	return plugin, db.SavePluginKind(plugin)
}

func FindAllPluginKind() (plugins []model.PluginType, err error) {

	return db.FindSpecifyAllPluginType(new(model.PluginType))
}
