package timego

import "time"

// Offset gets offset in seconds east of UTC from time.Time
func Offset(t time.Time) int {
	_, offset := t.Zone()

	return offset
}
