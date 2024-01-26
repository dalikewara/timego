package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFromString_OK tests FromString function with "OK" scenario
func TestFromString_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeString := "2023-12-26 15:40:20"
	timeNow, err := timego.FromString(timeString, timeLayout)

	assert.Nil(t, err)
	assert.Equal(t, false, timeNow.IsZero())
	assert.Equal(t, timeString, timeNow.Format(timeLayout))
}

// TestFromString_LayoutDoesntMatch tests FromString function with "LayoutDoesntMatch" scenario
func TestFromString_LayoutDoesntMatch(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05.999999999Z"
	timeString := "2023-12-26 15:40:20"
	timeNow, err := timego.FromString(timeString, timeLayout)

	assert.NotNil(t, err)
	assert.Equal(t, true, timeNow.IsZero())
}

// TestFromUnix_OK tests FromUnix function with "OK" scenario
func TestFromUnix_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeString := "2023-12-26 15:40:20"
	timeUnix := int64(1703605220)
	timeNow := timego.FromUnix(timeUnix)

	assert.Equal(t, false, timeNow.IsZero())
	assert.Equal(t, timeString, timeNow.Format(timeLayout))
}
