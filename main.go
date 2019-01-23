package main

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/GoTrigger/service"
	"github.com/andy-zhangtao/GoTrigger/trigger"
	"github.com/andy-zhangtao/GoTrigger/util"
	"github.com/sirupsen/logrus"
)

var _VERSION_, _BUILD_ string

func main() {
	logrus.WithFields(logrus.Fields{"VERSION": _VERSION_, "BUILD": _BUILD_}).Info(model.MODULENAME)
	util.InitTriggerChan()

	go trigger.QueryTrigger()
	service.GraphQL()
}
