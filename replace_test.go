package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestReplaceYear_OK tests ReplaceYear function with "OK" scenario
func TestReplaceYear_OK(t *testing.T) {
	timeNow := timego.ReplaceYear(time.Now(), 2024)

	assert.Equal(t, 2024, timeNow.Year())
}

// TestReplaceMonth_OK tests ReplaceMonth function with "OK" scenario
func TestReplaceMonth_OK(t *testing.T) {
	timeNow := timego.ReplaceMonth(time.Now(), 12)

	assert.Equal(t, time.Month(12), timeNow.Month())
}

// TestReplaceDay_OK tests ReplaceDay function with "OK" scenario
func TestReplaceDay_OK(t *testing.T) {
	timeNow := timego.ReplaceDay(time.Now(), 26)

	assert.Equal(t, 26, timeNow.Day())
}

// TestReplaceHour_OK tests ReplaceHour function with "OK" scenario
func TestReplaceHour_OK(t *testing.T) {
	timeNow := timego.ReplaceHour(time.Now(), 15)

	assert.Equal(t, 15, timeNow.Hour())
}

// TestReplaceMinute_OK tests ReplaceMinute function with "OK" scenario
func TestReplaceMinute_OK(t *testing.T) {
	timeNow := timego.ReplaceMinute(time.Now(), 40)

	assert.Equal(t, 40, timeNow.Minute())
}

// TestReplaceSecond_OK tests ReplaceSecond function with "OK" scenario
func TestReplaceSecond_OK(t *testing.T) {
	timeNow := timego.ReplaceSecond(time.Now(), 20)

	assert.Equal(t, 20, timeNow.Second())
}

// TestReplaceNanoSecond_OK tests ReplaceNanoSecond function with "OK" scenario
func TestReplaceNanoSecond_OK(t *testing.T) {
	timeNow := timego.ReplaceNanoSecond(time.Now(), 99)

	assert.Equal(t, 99, timeNow.Nanosecond())
}

// TestReplaceDate_OK tests ReplaceDate function with "OK" scenario
func TestReplaceDate_OK(t *testing.T) {
	timeNow := timego.ReplaceDate(time.Now(), 2024, 12, 26)

	assert.Equal(t, 2024, timeNow.Year())
	assert.Equal(t, time.Month(12), timeNow.Month())
	assert.Equal(t, 26, timeNow.Day())
}

// TestReplaceTime_OK tests ReplaceTime function with "OK" scenario
func TestReplaceTime_OK(t *testing.T) {
	timeNow := timego.ReplaceTime(time.Now(), 15, 40, 20, 99)

	assert.Equal(t, 15, timeNow.Hour())
	assert.Equal(t, 40, timeNow.Minute())
	assert.Equal(t, 20, timeNow.Second())
	assert.Equal(t, 99, timeNow.Nanosecond())
}
