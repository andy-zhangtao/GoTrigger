package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveTriggerPlugin(t *testing.T) {
	err := SaveTriggerPlugin(model.TriggerPlugin{
		ID:       bson.NewObjectId(),
		Name:     "plugin_first",
		Endpoint: "tcp://127.0.0.1",
		PID:      2,
		Desc:     "The first plugin",
	})

	assert.Nil(t, err)
}

func TestFindSpecifyTriggerPlugin(t *testing.T) {
	p := new(model.TriggerPlugin)
	p.Name = "plugin_first"
	err := FindSpecifyTriggerPlugin(p)
	assert.Nil(t, err)

	assert.Equal(t, "tcp://127.0.0.1", p.Endpoint)
	assert.Equal(t, 2, p.PID)
	assert.Equal(t, "The first plugin", p.Desc)
}

func TestFindSpecifyAllTriggerPlugin(t *testing.T) {
	ps, err := FindSpecifyAllTriggerPlugin(new(model.TriggerPlugin))
	assert.Nil(t, err)

	assert.Equal(t, 1, len(ps))
	assert.Equal(t, "tcp://127.0.0.1", ps[0].Endpoint)
	assert.Equal(t, 2, ps[0].PID)
	assert.Equal(t, "The first plugin", ps[0].Desc)
}

func TestUpdateSpecifyTriggerPlugin(t *testing.T) {
	p := new(model.TriggerPlugin)
	p.Name = "plugin_first"
	p.PID = 10
	p.Endpoint = "tcp://127.0.0.1:8000"
	p.Desc = "update"
	err := UpdateSpecifyTriggerPlugin(p, []string{"name"})
	assert.Nil(t, err)

	err = FindSpecifyTriggerPlugin(p)
	assert.Nil(t, err)

	assert.Equal(t, "tcp://127.0.0.1:8000", p.Endpoint)
	assert.Equal(t, 10, p.PID)
	assert.Equal(t, "update", p.Desc)
}
func TestDeleteTriggerPlugin(t *testing.T) {
	err := DeleteTriggerPlugin(&model.TriggerPlugin{
		Name: "plugin_first",
	})
	assert.Nil(t, err)

	ps, err := FindSpecifyAllTriggerPlugin(new(model.TriggerPlugin))
	assert.Nil(t, err)
	assert.Equal(t, 0, len(ps))
}
