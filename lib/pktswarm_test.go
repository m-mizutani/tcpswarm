package pktswarm_test

import (
	"testing"
	"time"

	pktswarm "github.com/m-mizutani/pktswarm/lib"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	sw, err := pktswarm.New(pktswarm.Config{
		FileName: "testdata/d1.pcap",
	})
	assert.Nil(t, err)

	ch, err := sw.Start()
	assert.Nil(t, err)
	recvCount := 0
	timeout := false

	select {
	case msg := <-ch:
		assert.NotEqual(t, 0, msg.Count)
		recvCount++
	case <-time.After(2 * time.Second):
		timeout = true
	}

	assert.False(t, timeout)
	assert.NotEqual(t, 0, recvCount)
}
