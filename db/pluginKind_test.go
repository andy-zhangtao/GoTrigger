package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSavePluginKind(t *testing.T) {
	err := SavePluginKind(model.PluginType{
		ID:   bson.NewObjectId(),
		Name: "http1",
		PID:  10,
		Desc: "A new http kind",
	})

	assert.Nil(t, err)

	err = SavePluginKind(model.PluginType{
		ID:   bson.NewObjectId(),
		Name: "http2",
		PID:  20,
		Desc: "A new http kind",
	})

	assert.Nil(t, err)
}

func TestFindSpecifyPluginType(t *testing.T) {
	p := model.PluginType{
		PID: 10,
	}
	err := FindSpecifyPluginType(&p)
	assert.Nil(t, err)

	assert.Equal(t, "http1", p.Name)
}

func TestUpdatePluginType(t *testing.T) {
	k := new(model.PluginType)

	k.PID = 10
	k.Name = "newHttp1"
	UpdatePluginType(k, []string{"pid"})

	p := model.PluginType{
		PID: 10,
	}
	err := FindSpecifyPluginType(&p)
	assert.Nil(t, err)

	assert.Equal(t, "newHttp1", p.Name)
}

func TestDeleteAllPluginType(t *testing.T) {
	DeleteAllPluginType()

	ks, err := FindSpecifyAllPluginType(new(model.PluginType))

	assert.Nil(t, err)
	assert.Equal(t, 0, len(ks))
}
