package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestAddYears_OK tests AddYears function with "OK" scenario
func TestAddYears_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddYears(timeNow, 1)

	assert.Equal(t, timeNow.Year()+1, timeAdd.Year())
}

// TestAddMonths_OK tests AddMonths function with "OK" scenario
func TestAddMonths_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddMonths(timeNow, 1)

	assert.Equal(t, timeNow.Month()+1, timeAdd.Month())
}

// TestAddDays_OK tests AddDays function with "OK" scenario
func TestAddDays_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddDays(timeNow, 1)

	assert.Equal(t, timeNow.Day()+1, timeAdd.Day())
}

// TestAddHours_OK tests AddHours function with "OK" scenario
func TestAddHours_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddHours(timeNow, 1)

	assert.Equal(t, timeNow.Hour()+1, timeAdd.Hour())
}

// TestAddMinutes_OK tests AddMinutes function with "OK" scenario
func TestAddMinutes_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddMinutes(timeNow, 1)

	assert.Equal(t, timeNow.Minute()+1, timeAdd.Minute())
}

// TestAddSeconds_OK tests AddSeconds function with "OK" scenario
func TestAddSeconds_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddSeconds(timeNow, 1)

	assert.Equal(t, timeNow.Second()+1, timeAdd.Second())
}

// TestAddNanoSeconds_OK tests AddNanoSeconds function with "OK" scenario
func TestAddNanoSeconds_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddNanoSeconds(timeNow, 1)

	assert.Equal(t, timeNow.Nanosecond()+1, timeAdd.Nanosecond())
}

// TestAddDuration_OK tests AddDuration function with "OK" scenario
func TestAddDuration_OK(t *testing.T) {
	timeNow := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	timeAdd := timego.AddDuration(timeNow, 5*time.Minute)

	assert.Equal(t, timeNow.Minute()+5, timeAdd.Minute())
}
