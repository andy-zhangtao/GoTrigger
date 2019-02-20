package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
)

func SaveTriggerJnl(t model.Trigger, status int, message string) error {
	jnl := model.TriggerJnl{
		Name:    t.Name,
		Status:  status,
		Message: message,
	}

	return db.SaveTriggerJnl(jnl)
}

func QueryAllTriggerJnl() (jnl []model.TriggerJnl, err error) {

	return db.FetchAllTriggerJnl()
}

func DeleteALlTriggerJnl() (err error) {
	return db.DeleteAllTriggerJnl(new(model.TriggerJnl))
}
