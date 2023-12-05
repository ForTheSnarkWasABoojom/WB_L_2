package main

import (
	"testing"

	"github.com/beevik/ntp"
	"github.com/stretchr/testify/assert"
)

func TestGetNTPTime(t *testing.T) {
	ntpTime, err := ntp.Time("pool.ntp.org")
	assert.NoError(t, err)
	assert.NotNil(t, ntpTime)
}
