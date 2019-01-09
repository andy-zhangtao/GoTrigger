package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/bwidow"
	"github.com/sirupsen/logrus"
)

//var endpoint = os.Getenv(model.ENV_AGENT_MONGO_ENDPOINT)
//var username = os.Getenv(model.ENV_AGENT_MONGO_NAME)
//var password = os.Getenv(model.ENV_AGENT_MONGO_PASSWD)
//var dbname = os.Getenv(model.ENV_AGENT_MONGO_DBNAME)
//var session *mgo.Session
var bw *bwidow.BW

func init() {

	bw = bwidow.GetWidow()
	err := bw.Driver(bwidow.DRIVER_MONGO).Error()
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{"Connect Blank Widow Success BW Version": bw.Version()}).Info(model.MODULENAME)

	bw = bw.Map(model.Trigger{}, model.DB_TRIGGER).Map(model.TriggerType{}, model.DB_TRIGGER_TYPE)

	if err := bw.CheckIndex(new(model.Trigger)).Error(); err != nil {
		logrus.WithFields(logrus.Fields{"CheckIndexError": err}).Error(model.MODULENAME)
		logrus.Panic(err)
	}

	logrus.WithFields(logrus.Fields{"BWidow Init Map": "Success", "Check Index": "Sucess"}).Info(model.MODULENAME)
}
