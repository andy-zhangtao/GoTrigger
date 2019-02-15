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

	var triggers []*model.Trigger

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
		triggers = append(triggers, &t)
	}

	var ticker = time.NewTicker(tick * time.Second)
	now := time.Now().Unix()

	for {
		select {
		case <-ticker.C:
			now += tick
			for _, t := range triggers {
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
		case id := <-util.GetTriggerChan():
			t := model.Trigger{
				ID: id,
			}
			err := db.FindSpecifyTrigger(&t)
			if err != nil {
				logrus.WithFields(logrus.Fields{"Query-Trigger-Error": err}).Error(model.MODULENAME)
				continue
			}

			triggers = append(triggers, &t)
		}
	}
}
