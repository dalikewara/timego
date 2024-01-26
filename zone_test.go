package timego_test

import (
	"github.com/dalikewara/timego"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestZone_OK tests Zone function with "OK" scenario
func TestZone_OK(t *testing.T) {
	assert.NotEmpty(t, timego.Zone(time.Now()))
}
