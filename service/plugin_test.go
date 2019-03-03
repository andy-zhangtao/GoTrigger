package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanPlugin(t *testing.T) {
	db.DeleteALlTriggerPlugin(new(model.TriggerPlugin))
}

func TestAddNewPlugin(t *testing.T) {
	_, err := AddNewPlugin(model.TriggerPlugin{
		ID:       bson.NewObjectId(),
		Name:     "plugin1",
		Endpoint: "127.0.0.1",
		PID:      10,
		Desc:     "from unit test",
	})

	assert.Nil(t, err)

	_, err = AddNewPlugin(model.TriggerPlugin{
		ID:       bson.NewObjectId(),
		Name:     "plugin2",
		Endpoint: "127.0.1.1",
		PID:      20,
		Desc:     "from unit test",
	})

	assert.Nil(t, err)
}

func TestFindSpecifyPlugin(t *testing.T) {
	p, err := FindSpecifyPlugin(10)
	assert.Nil(t, err)

	assert.Equal(t, "plugin1", p.Name)
	assert.Equal(t, "127.0.0.1", p.Endpoint)
	assert.Equal(t, "from unit test", p.Desc)
}

func TestFindAllPlugin(t *testing.T) {
	ps, err := FindAllPlugin()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(ps))

	if ps[0].PID == 10 {
		assert.Equal(t, "plugin1", ps[0].Name)
		assert.Equal(t, "127.0.0.1", ps[0].Endpoint)
		assert.Equal(t, "from unit test", ps[0].Desc)
	} else {
		assert.Equal(t, "plugin2", ps[1].Name)
		assert.Equal(t, "127.0.1.1", ps[1].Endpoint)
		assert.Equal(t, "from unit test", ps[1].Desc)
	}
}

func TestUpdateSpecifyPlugin(t *testing.T) {
	err := UpdateSpecifyPlugin(model.TriggerPlugin{
		Name:     "plugin1",
		Endpoint: "192.168.0.1",
	})

	assert.Nil(t, err)

	p, err := FindSpecifyPlugin(10)
	assert.Nil(t, err)

	assert.Equal(t, "plugin1", p.Name)
	assert.Equal(t, "192.168.0.1", p.Endpoint)
	assert.Equal(t, "from unit test", p.Desc)
}

func TestDeleteSpecifyPlugin(t *testing.T) {
	err := DeleteSpecifyPlugin("plugin1")
	assert.Nil(t, err)

	err = DeleteSpecifyPlugin("plugin2")
	assert.Nil(t, err)

	ps, err := FindAllPlugin()
	assert.Nil(t, err)

	assert.Equal(t, 0, len(ps))
}
