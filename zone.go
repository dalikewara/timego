package timego

import "time"

// Zone gets abbreviated name of the time zone from time.Time (such as "WIB", "CET", etc)
func Zone(t time.Time) string {
	zone, _ := t.Zone()

	return zone
}
