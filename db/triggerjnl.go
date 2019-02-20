package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	zt "github.com/andy-zhangtao/gogather/time"
	"github.com/globalsign/mgo/bson"
)

func SaveTriggerJnl(jnl model.TriggerJnl) error {
	if jnl.ID == "" {
		jnl.ID = bson.NewObjectId()
	}

	t := zt.Ztime{}
	jnl.Time, _ = t.Now().Format("YYYYMMDD hh:mm:ss")

	return bw.Save(jnl)
}

func FetchAllTriggerJnl() (allTriggerJnl []model.TriggerJnl, err error) {
	err = bw.FindAllWithSort(new(model.TriggerJnl), &allTriggerJnl, []string{"-time"})
	return
}

func DeleteAllTriggerJnl(jnlPtr *model.TriggerJnl) (err error) {
	_, err = bw.DeleteAll(jnlPtr)
	return
}
