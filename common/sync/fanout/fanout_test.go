package fanout

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFanout_Close(t *testing.T) {
	fan := New(Worker(1), Buffer(1024))
	fan.Close()
	assert.NotNil(t, fan.Do(context.Background(), func(c context.Context) {}))
}

func TestFanout_Do(t *testing.T) {
	fan := New(Worker(1), Buffer(1024))
	var run bool
	assert.Nil(t, fan.Do(context.Background(), func(c context.Context) {
		run = true
		panic("error")
	}))
	time.Sleep(50 * time.Millisecond)
	assert.True(t, run)
}
