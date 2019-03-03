package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/sirupsen/logrus"
)

func GetNextSeq() (seq int) {

	s := new(model.Sequence)
	err := bw.FindOne(s)
	if err != nil {
		return 0
	}

	s.SID += 1

	seq = s.SID

	if _, err := bw.Update(s, []string{"desc"}); err != nil {
		logrus.WithFields(logrus.Fields{"update sequence error": err}).Error(model.MODULENAME)
	}

	logrus.WithFields(logrus.Fields{"get sequence ": seq}).Error(model.MODULENAME)
	return
}
