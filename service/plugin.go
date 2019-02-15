package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
)

func AddNewPlugin(plugin model.TriggerPlugin) (model.TriggerPlugin, error) {
	plugin.ID = bson.NewObjectId()
	return plugin, db.SaveTriggerPlugin(plugin)
}

func FindSpecifyPlugin(pid int) (plugin model.TriggerPlugin, err error) {
	plugin.PID = pid

	err = db.FindSpecifyTriggerPlugin(&plugin)
	return
}

func FindAllPlugin() (plugins []model.TriggerPlugin, err error) {
	return db.FindSpecifyAllTriggerPlugin(new(model.TriggerPlugin))
}

func DeleteSpecifyPlugin(name string) (err error) {
	return db.DeleteTriggerPlugin(&model.TriggerPlugin{
		Name: name,
	})
}
