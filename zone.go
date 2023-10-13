package timego

import "time"

func Zone(t time.Time) string {
	zone, _ := t.Zone()

	return zone
}
