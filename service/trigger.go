package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
)

func AddNewTrigger(trigger model.Trigger) error {
	trigger.ID = bson.NewObjectId()
	return db.SaveTrigger(trigger)
}

func UpdateTriggerNextTime(name string, next uint64) error {
	t := model.Trigger{
		Name:     name,
		NextTime: next,
	}

	return db.UpdateTrigger(&t, []string{
		"name",
	})
}

func DisableTrigger(name string) error {
	t := model.Trigger{
		Name:   name,
		Enable: false,
	}

	return db.UpdateTrigger(&t, []string{
		"name",
	})
}

func EnableTrigger(name string) error {
	t := model.Trigger{
		Name:   name,
		Enable: true,
	}

	return db.UpdateTrigger(&t, []string{
		"name",
	})
}

func FindSpecifyTrigger(name string) (trigger model.Trigger, err error) {
	trigger.Name = name
	err = db.FindSpecifyTrigger(&trigger)
	return
}

func FindSpecifyAllTrigger() (ts []model.Trigger, err error) {
	return db.FindSpecifyAllTrigger(new(model.Trigger))
}
