package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanTrigger(t *testing.T) {
	DeleTrigger("trigger")
}

func TestAddNewTrigger(t *testing.T) {
	_, err := AddNewTrigger(model.Trigger{
		Name:     "trigger",
		Async:    true,
		Enable:   false,
		Interval: 200,
		Parallel: 20,
		Desc:     "unit",
	})

	assert.Nil(t, err)
}

func TestEnableTrigger(t *testing.T) {
	err := EnableTrigger("trigger")

	assert.Nil(t, err)

	tr, err := FindSpecifyTrigger("trigger")
	assert.Nil(t, err)

	assert.Equal(t, true, tr.Enable)
	assert.Equal(t, true, tr.Async)
	assert.Equal(t, 200, tr.Interval)
	assert.Equal(t, 20, tr.Parallel)
	assert.Equal(t, "unit", tr.Desc)
}

func TestDisableTrigger(t *testing.T) {
	err := DisableTrigger("trigger")
	assert.Nil(t, err)

	tr, err := FindSpecifyTrigger("trigger")
	assert.Nil(t, err)

	assert.Equal(t, false, tr.Enable)
	assert.Equal(t, true, tr.Async)
	assert.Equal(t, 200, tr.Interval)
	assert.Equal(t, 20, tr.Parallel)
	assert.Equal(t, "unit", tr.Desc)
}

func TestFindSpecifyTrigger(t *testing.T) {
	tr, err := FindSpecifyTrigger("trigger")
	assert.Nil(t, err)

	assert.Equal(t, false, tr.Enable)
	assert.Equal(t, true, tr.Async)
	assert.Equal(t, 200, tr.Interval)
	assert.Equal(t, 20, tr.Parallel)
	assert.Equal(t, "unit", tr.Desc)
}

func TestDeleTrigger(t *testing.T) {
	err := DeleTrigger("trigger")
	assert.Nil(t, err)

	_, err = FindSpecifyTrigger("trigger")
	assert.Equal(t, "not found", err.Error())
}
