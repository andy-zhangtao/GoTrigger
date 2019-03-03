package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInterval(t *testing.T) {
	interval := "*:*:00"
	dt, err := ParseInterval(interval)
	assert.Error(t, err)
	assert.Equal(t, 0, dt)

	interval = "*:*:01"
	dt, err = ParseInterval(interval)
	assert.Error(t, err)
	assert.Equal(t, 0, dt)

	interval = "*-*-* *:*:01"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 1, dt)

	interval = "*-*-* *:*:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 30, dt)

	interval = "*-*-* *:*:130"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 60, dt)

	interval = "*-*-* *:01:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 90, dt)

	interval = "*-*-* *:01:130"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 120, dt)

	interval = "*-*-* *:60:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 60*60+30, dt)

	interval = "*-*-* *:70:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 60*60+30, dt)

	interval = "*-*-* 01:20:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 1*(60*60)+20*60+30, dt)

	interval = "*-*-* 00:20:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 0*(60*60)+20*60+30, dt)

	interval = "*-*-* 23:20:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 23*(60*60)+20*60+30, dt)

	interval = "*-*-* 25:20:30"
	dt, err = ParseInterval(interval)
	assert.Empty(t, err)
	assert.Equal(t, 24*(60*60)+20*60+30, dt)
}
