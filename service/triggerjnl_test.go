package service

import (
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClean(t *testing.T) {
	db.DeleteAllTriggerJnl(new(model.TriggerJnl))
}

func TestSaveTriggerJnl(t *testing.T) {
	err := SaveTriggerJnl(model.Trigger{
		Name: "http-trigger-1",
	}, 1, "success")

	assert.Nil(t, err)

	err = SaveTriggerJnl(model.Trigger{
		Name: "http-trigger-2",
	}, 2, "failed")

	assert.Nil(t, err)
}

func TestQueryAllTriggerJnl(t *testing.T) {
	jnl, err := QueryAllTriggerJnl()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(jnl))

	if jnl[0].Name == "http-trigger-1" {
		assert.Equal(t, 1, jnl[0].Status)
		assert.Equal(t, "success", jnl[0].Message)
	} else {
		assert.Equal(t, 2, jnl[1].Status)
		assert.Equal(t, "failed", jnl[1].Message)
	}

}
