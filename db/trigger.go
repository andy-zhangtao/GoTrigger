package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	zt "github.com/andy-zhangtao/gogather/time"
)

func SaveTrigger(trigger model.Trigger) error {
	t := zt.Ztime{}
	trigger.CreateTime, _ = t.Now().Format("YYYYMMDD hh:mm:ss")

	return bw.Save(trigger)
}

func UpdateTrigger(tPtr *model.Trigger, filter []string) (err error) {
	t := zt.Ztime{}
	tPtr.UpdateTime, _ = t.Now().Format("YYYYMMDD hh:mm:ss")
	_, err = bw.Update(tPtr, filter)

	return
}

func DeleteTrigger(tPtr *model.Trigger, filter []string) (err error) {
	_, err = bw.Delete(tPtr, filter)
	return
}

func DeleteAllTrigger(tPtr *model.Trigger) (err error) {
	_, err = bw.DeleteAll(tPtr)
	return
}

func FindSpecifyTrigger(tPtr *model.Trigger, filter []string) (err error) {
	return bw.FindOne(tPtr, filter...)
}

func FindSpecifyAllTrigger(tPtr *model.Trigger) (allTriggers []model.Trigger, err error) {
	err = bw.FindAll(tPtr, &allTriggers, "async", "enable")

	return
}
