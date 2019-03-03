package db

import "github.com/andy-zhangtao/GoTrigger/model"

func SaveTriggerPlugin(p model.TriggerPlugin) error {
	return bw.Save(p)
}

func DeleteTriggerPlugin(pPtr *model.TriggerPlugin) (err error) {
	_, err = bw.Delete(pPtr, []string{"name"})
	return
}

func DeleteALlTriggerPlugin(pPtr *model.TriggerPlugin) (err error) {
	_, err = bw.DeleteAll(pPtr)
	return
}

func FindSpecifyTriggerPlugin(pPtr *model.TriggerPlugin) (err error) {
	return bw.FindOne(pPtr)
}

func FindSpecifyAllTriggerPlugin(pPtr *model.TriggerPlugin) (allPlugin []model.TriggerPlugin, err error) {
	err = bw.FindAll(pPtr, &allPlugin)
	return
}

func UpdateSpecifyTriggerPlugin(pPtr *model.TriggerPlugin, filter []string) (err error) {
	_, err = bw.Update(pPtr, filter)
	return
}
