package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

func AddNewTrigger(trigger model.Trigger) (model.Trigger, error) {
	trigger.ID = bson.NewObjectId()
	return trigger, db.SaveTrigger(trigger)
}

func DeleTrigger(name string) error {
	return db.DeleteTrigger(&model.Trigger{
		Name: name,
	}, []string{"name"})
}

func UpdateTrigger(trigger model.Trigger) (err error) {
	return db.UpdateTrigger(&trigger, []string{
		"name",
	}, "enable", "async")
}

func UpdateTriggerNextTime(name string, next int64) error {
	t := model.Trigger{
		Name:     name,
		NextTime: next,
	}

	logrus.WithFields(logrus.Fields{"name": t.Name, "nextime": t.NextTime}).Info(model.MODULENAME)
	return db.UpdateTrigger(&t, []string{
		"name",
	}, "enable", "async")
}

func DisableTrigger(name string) error {
	t := model.Trigger{
		Name:   name,
		Enable: false,
	}

	//logrus.WithFields(logrus.Fields{"trigger": t}).Info(model.MODULENAME)
	return db.UpdateTrigger(&t, []string{
		"name",
	}, "async")
}

func EnableTrigger(name string) error {
	t := model.Trigger{
		Name:   name,
		Enable: true,
	}

	return db.UpdateTrigger(&t, []string{
		"name",
	}, "async")
}

func FindSpecifyTrigger(name string) (trigger model.Trigger, err error) {
	trigger.Name = name
	err = db.FindSpecifyTrigger(&trigger, []string{"name"})
	return
}

func FindSpecifyAllTrigger() (ts []model.Trigger, err error) {
	return db.FindSpecifyAllTrigger(new(model.Trigger))
}
