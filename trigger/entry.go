package trigger

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/GoTrigger/service"
	"github.com/andy-zhangtao/GoTrigger/util"
	"github.com/sirupsen/logrus"
	"time"
)

const tick = 3

func QueryTrigger() error {

	triggers := make(map[*model.Trigger]bool)

	t := model.Trigger{
		Enable: true,
	}
	_triggers, err := db.FindSpecifyAllTrigger(&t)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Query-All-Trigger-Error": err}).Error(model.MODULENAME)
		return err
	}

	for _, t := range _triggers {
		logrus.WithFields(logrus.Fields{"name": t.Name, "parallel": t.Parallel}).Info(model.MODULENAME)
		triggers[&t] = t.Enable
	}

	var ticker = time.NewTicker(tick * time.Second)
	now := time.Now().Unix()

	for {
		select {
		case <-ticker.C:
			now += tick
			for t, enable := range triggers {
				logrus.WithFields(logrus.Fields{"name": t.Name, "enable": t.Enable}).Info(model.MODULENAME)
				if enable {
					if now >= t.NextTime {
						t.NextTime = now + int64(t.Interval)
						if err := execut(t); err != nil {
							logrus.WithFields(logrus.Fields{"trigger-error": err}).Error(model.MODULENAME)
						}
						if err := service.UpdateTriggerNextTime(t.Name, t.NextTime); err != nil {
							logrus.WithFields(logrus.Fields{"update-next-time-error": err}).Error(model.MODULENAME)
						}
					}
				}
			}
		case id := <-util.GetTriggerChan():
			logrus.WithFields(logrus.Fields{"Action": "Reload"}).Info(model.MODULENAME)
			t := model.Trigger{
				ID: id,
			}
			err := db.FindSpecifyTrigger(&t, []string{"_id"})
			if err != nil {
				logrus.WithFields(logrus.Fields{"Query-Trigger-Error": err}).Error(model.MODULENAME)
				continue
			}

			logrus.WithFields(logrus.Fields{"new trigger": t}).Info(model.MODULENAME)
			for _t, _ := range triggers {
				if _t.ID == t.ID {
					//更新旧数据
					delete(triggers, _t)
					break
				}
			}

			triggers[&t] = t.Enable

		case t := <-util.GetFireChan():
			if err := execut(&t); err != nil {
				logrus.WithFields(logrus.Fields{"trigger-error": err}).Error(model.MODULENAME)
			}
		}
	}
}
