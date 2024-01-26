package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestOffset_OK tests Offset function with "OK" scenario
func TestOffset_OK(t *testing.T) {
	zone, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().UTC().In(zone)
	offset := timego.Offset(timeNow)

	assert.NotZero(t, offset)
	assert.Equal(t, 25200, offset)
}
