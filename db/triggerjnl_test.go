package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveTriggerJnl(t *testing.T) {
	DeleteAllTriggerJnl(new(model.TriggerJnl))

	err := SaveTriggerJnl(model.TriggerJnl{
		ID:     bson.NewObjectId(),
		Name:   "first",
		Status: 2,
	})

	err = SaveTriggerJnl(model.TriggerJnl{
		ID:     bson.NewObjectId(),
		Name:   "second",
		Status: 3,
	})

	assert.Nil(t, err)
}

func TestFetchAllTriggerJnl(t *testing.T) {
	fets, err := FetchAllTriggerJnl()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(fets))

	for _, f := range fets {
		if f.Name == "first" {
			assert.Equal(t, 2, f.Status)
		} else {
			assert.Equal(t, 3, f.Status)
		}
	}
}
