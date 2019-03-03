package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanPluginType(t *testing.T) {
	DeleteAllPluginKind()
}

func TestAddNewPluginKind(t *testing.T) {
	_, err := AddNewPluginKind(model.PluginType{
		ID:   bson.NewObjectId(),
		Name: "http1",
		PID:  10,
		Desc: "A new http kind",
	})

	assert.Nil(t, err)

	_, err = AddNewPluginKind(model.PluginType{
		ID:   bson.NewObjectId(),
		Name: "http2",
		PID:  20,
		Desc: "A new http kind 2",
	})

	assert.Nil(t, err)
}

func TestFindSpecifyPluginKind(t *testing.T) {
	kind, err := FindSpecifyPluginKind(10)
	assert.Nil(t, err)

	assert.Equal(t, "http1", kind.Name)
	assert.Equal(t, "A new http kind", kind.Desc)
}

func TestUpdatePluginKind(t *testing.T) {
	err := UpdateSpecifyPluginKind(model.PluginType{
		PID:  20,
		Name: "newHttp2",
	})

	assert.Nil(t, err)
}

func TestFindAllPluginKind(t *testing.T) {
	ks, err := FindAllPluginKind()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(ks))
	if ks[0].PID == 10 {
		assert.Equal(t, "http1", ks[0].Name)
		assert.Equal(t, "A new http kind", ks[0].Desc)
	} else {
		assert.Equal(t, "newHttp2", ks[1].Name)
		assert.Equal(t, "A new http kind 2", ks[1].Desc)
	}
}

func TestDeleteAllPluginKind(t *testing.T) {
	err := DeleteAllPluginKind()
	assert.Nil(t, err)
	ks, err := FindAllPluginKind()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(ks))
}
