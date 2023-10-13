package timego

import "time"

func Offset(t time.Time) int {
	_, offset := t.Zone()

	return offset
}
