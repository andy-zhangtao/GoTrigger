package db

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveTrigger(t *testing.T) {

	DeleteAllTrigger(new(model.Trigger))
	
	tr := model.Trigger{
		ID:       bson.NewObjectId(),
		Name:     "first",
		Enable:   true,
		Parallel: 10,
	}
	err := SaveTrigger(tr)
	assert.Nil(t, err)

	tr = model.Trigger{
		ID:       bson.NewObjectId(),
		Name:     "second",
		Enable:   true,
		Parallel: 2,
	}
	err = SaveTrigger(tr)
	assert.Nil(t, err)
}

func TestFindSpecifyTrigger(t *testing.T) {
	tr := new(model.Trigger)
	tr.Name = "first"
	err := FindSpecifyTrigger(tr, []string{"name"})
	assert.Nil(t, err)

	assert.Equal(t, true, tr.Enable)
	assert.Equal(t, 10, tr.Parallel)

	tr.Name = "second"
	err = FindSpecifyTrigger(tr, []string{"name"})
	assert.Nil(t, err)

	assert.Equal(t, false, tr.Async)
	assert.Equal(t, 2, tr.Parallel)
}

func TestFindSpecifyAllTrigger(t *testing.T) {
	trs, err := FindSpecifyAllTrigger(new(model.Trigger))
	assert.Nil(t, err)

	assert.Equal(t, 2, len(trs))

	if trs[0].Name == "first" {
		assert.Equal(t, 10, trs[0].Parallel)
	} else {
		assert.Equal(t, 2, trs[1].Parallel)
	}
}

func TestUpdateTrigger(t *testing.T) {
	tr := new(model.Trigger)
	tr.Name = "first"
	tr.Enable = false
	tr.Async = true
	tr.Parallel = 22
	tr.Interval = 9
	err := UpdateTrigger(tr, []string{"name"})
	assert.Nil(t, err)
}

func TestFindSpecifyTriggerAfterUpdate(t *testing.T) {
	tr := new(model.Trigger)
	tr.Name = "first"
	err := FindSpecifyTrigger(tr, []string{"name"})
	assert.Nil(t, err)

	assert.Equal(t, false, tr.Enable)
	assert.Equal(t, true, tr.Async)
	assert.Equal(t, 22, tr.Parallel)
	assert.Equal(t, 9, tr.Interval)

	tr.Name = "second"
	err = FindSpecifyTrigger(tr, []string{"name"})
	assert.Nil(t, err)

	assert.Equal(t, false, tr.Async)
	assert.Equal(t, 2, tr.Parallel)
}

func TestDeleteTrigger(t *testing.T) {
	tr := new(model.Trigger)
	tr.Name = "first"
	err := DeleteTrigger(tr, []string{"name"})
	assert.Nil(t, err)
}

func TestFindSpecifyAllTriggerAfterDelete(t *testing.T) {
	trs, err := FindSpecifyAllTrigger(new(model.Trigger))
	assert.Nil(t, err)

	assert.Equal(t, 1, len(trs))
}

func TestDeleteAllTrigger(t *testing.T) {
	DeleteAllTrigger(new(model.Trigger))
	trs, err := FindSpecifyAllTrigger(new(model.Trigger))
	assert.Nil(t, err)

	assert.Equal(t, 0, len(trs))
}
