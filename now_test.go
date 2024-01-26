package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNow_OK tests Now function with "OK" scenario
func TestNow_OK(t *testing.T) {
	timeNow := timego.Now()

	assert.Equal(t, false, timeNow.IsZero())
}

// TestNowUTC_OK tests NowUTC function with "OK" scenario
func TestNowUTC_OK(t *testing.T) {
	timeNow := timego.NowUTC()

	assert.Equal(t, false, timeNow.IsZero())
}
