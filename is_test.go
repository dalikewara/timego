package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestIsExpired_OK tests IsExpired function with "OK" scenario
func TestIsExpired_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeExpired := time.Date(2023, 1, 1, 1, 5, 0, 0, time.UTC)
	isExpired := timego.IsExpired(timeNow, timeExpired)

	assert.Equal(t, false, isExpired)

	timeNow = time.Date(2023, 1, 1, 1, 6, 0, 0, time.UTC)
	timeExpired = time.Date(2023, 1, 1, 1, 5, 0, 0, time.UTC)
	isExpired = timego.IsExpired(timeNow, timeExpired)

	assert.Equal(t, true, isExpired)
}
