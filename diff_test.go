package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestDiff_OK tests Diff function with "OK" scenario
func TestDiff_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeOldString := "2023-12-25 14:40:20"
	timeCurrentString := "2023-12-26 15:40:20"
	timeOld, _ := time.Parse(timeLayout, timeOldString)
	timeCurrent, _ := time.Parse(timeLayout, timeCurrentString)
	timeDiff := timego.Diff(timeOld, timeCurrent)

	assert.Equal(t, float64(25), timeDiff.Hours())
}

// TestDiffInDays_OK tests DiffInDays function with "OK" scenario
func TestDiffInDays_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeOldString := "2023-12-25 14:40:20"
	timeCurrentString := "2023-12-26 15:40:20"
	timeOld, _ := time.Parse(timeLayout, timeOldString)
	timeCurrent, _ := time.Parse(timeLayout, timeCurrentString)
	timeDiff := timego.DiffInDays(timeOld, timeCurrent)

	assert.Equal(t, 1, timeDiff)
}

// TestDiffInSeconds_OK tests DiffInSeconds function with "OK" scenario
func TestDiffInSeconds_OK(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeOldString := "2023-12-25 14:40:20"
	timeCurrentString := "2023-12-26 15:40:20"
	timeOld, _ := time.Parse(timeLayout, timeOldString)
	timeCurrent, _ := time.Parse(timeLayout, timeCurrentString)
	timeDiff := timego.DiffInSeconds(timeOld, timeCurrent)

	assert.Equal(t, float64(90000), timeDiff)
}
