package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestToString_OK tests ToString function with "OK" scenario
func TestToString_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStr := timego.ToString(timeNow, timeLayout)

	assert.NotEmpty(t, timeStr)
	assert.Equal(t, 19, len(timeStr))
	assert.Equal(t, "2023-01-01 01:00:00", timeStr)
}

// TestToString_OK tests ToString function with "InvalidTimeLayout" scenario
func TestToString_InvalidTimeLayout(t *testing.T) {
	t.Run("InvalidTimeLayout_1", func(t *testing.T) {
		timeLayout := "%*&^*"
		timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
		timeStr := timego.ToString(timeNow, timeLayout)

		assert.NotEmpty(t, timeStr)
		assert.Equal(t, len(timeLayout), len(timeStr))
		assert.Equal(t, timeLayout, timeStr)
	})

	t.Run("InvalidTimeLayout_2", func(t *testing.T) {
		timeLayout := "2006-01-X 15:04:05"
		timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
		timeStr := timego.ToString(timeNow, timeLayout)

		assert.NotEmpty(t, timeStr)
		assert.Equal(t, 18, len(timeStr))
		assert.Equal(t, "2023-01-X 01:00:00", timeStr)
	})
}

// TestToUTC_OK tests ToUTC function with "OK" scenario
func TestToUTC_OK(t *testing.T) {
	timeUTC := timego.ToUTC(time.Now())
	zone, _ := timeUTC.Zone()

	assert.Equal(t, "UTC", zone)
}

// TestToLocation_OK tests ToLocation function with "OK" scenario
func TestToLocation_OK(t *testing.T) {
	location := "America/Los_Angeles"
	timeNow, err := timego.ToLocation(time.Now(), location)

	assert.Nil(t, err)
	assert.Equal(t, false, timeNow.IsZero())
	assert.Equal(t, location, timeNow.Location().String())
}

// TestToLocation_InvalidLocation tests ToLocation function with "InvalidLocation" scenario
func TestToLocation_InvalidLocation(t *testing.T) {
	location := "America/Los_Angeles_Invalid"
	timeNow, err := timego.ToLocation(time.Now(), location)

	assert.NotNil(t, err)
	assert.Equal(t, false, timeNow.IsZero())
}
