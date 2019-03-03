package db

import "github.com/andy-zhangtao/GoTrigger/model"

func SavePluginKind(p model.PluginType) error {
	return bw.Save(p)
}

func DeletePluginType(pPtr *model.PluginType) (err error) {
	_, err = bw.DeleteAll(pPtr)
	return
}

func FindSpecifyPluginType(pPtr *model.PluginType) (err error) {
	return bw.FindOne(pPtr, "pid")
}

func FindSpecifyAllPluginType(pPtr *model.PluginType) (allPlugin []model.PluginType, err error) {
	err = bw.FindAll(pPtr, &allPlugin)
	return
}

func DeleteAllPluginType() (err error) {
	_, err = bw.DeleteAll(new(model.PluginType))
	return
}

func UpdatePluginType(kPtr *model.PluginType, filter []string) (err error) {
	_, err = bw.Update(kPtr, filter)
	return
}