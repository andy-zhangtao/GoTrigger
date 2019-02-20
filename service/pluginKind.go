package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
)

func AddNewPluginKind(plugin model.PluginType) (model.PluginType, error) {
	plugin.ID = bson.NewObjectId()
	if plugin.PID == 0 {
		plugin.PID = db.GetNextSeq()
	}

	return plugin, db.SavePluginKind(plugin)
}

func FindAllPluginKind() (plugins []model.PluginType, err error) {

	return db.FindSpecifyAllPluginType(new(model.PluginType))
}

func FindSpecifyPluginKind(id int) (kind model.PluginType, err error) {
	kind.PID = id
	err = db.FindSpecifyPluginType(&kind)
	return
}

func DeleteAllPluginKind() (err error) {
	return db.DeleteAllPluginType()
}
