package db

import (
	"errors"
	"fmt"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/bwidow"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var endpoint = os.Getenv(model.ENV_AGENT_MONGO_ENDPOINT)
var username = os.Getenv(model.ENV_AGENT_MONGO_NAME)
var password = os.Getenv(model.ENV_AGENT_MONGO_PASSWD)
var dbname = os.Getenv(model.ENV_AGENT_MONGO_DBNAME)
var session *mgo.Session
var bw *bwidow.BW

func check() error {
	if endpoint == "" {
		return errors.New(fmt.Sprintf("[%s] Not Found", model.ENV_AGENT_MONGO_ENDPOINT))
	}

	if dbname == "" {
		return errors.New(fmt.Sprintf("[%s] Not Found", model.ENV_AGENT_MONGO_DBNAME))
	}
	return nil
}

func init() {
	err := check()
	if err != nil {
		logrus.Panic(err)
	}

	logrus.WithFields(logrus.Fields{"Connect Mongo": endpoint}).Info(model.MODULENAME)
	if username != "" || password != "" {
		dialInfo := &mgo.DialInfo{
			Addrs:    []string{endpoint},
			Database: dbname,
			Username: username,
			Password: password,
			Timeout:  10 * time.Second,
		}

		session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(err)
		}
	} else {
		session, err = mgo.Dial(endpoint)
	}
	b, err := session.BuildInfo()
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{"Mongo Server": b.Version}).Info(model.MODULENAME)

	bw = bwidow.GetWidow()
	err = bw.Driver(bwidow.DRIVER_MONGO).Error()
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{"Connect Blank Widow Success BW Version": bw.Version()}).Info(model.MODULENAME)

	bw = bw.Map(model.Trigger{}, model.DB_TRIGGER).Map(model.TriggerType{}, model.DB_TRIGGER_TYPE)

	if err := bw.CheckIndex(new(model.Trigger)); err != nil {
		logrus.Panic(err)
	}

	logrus.WithFields(logrus.Fields{"BWidow Init Map": "Success", "Check Index": "Sucess"}).Info(model.MODULENAME)
}
