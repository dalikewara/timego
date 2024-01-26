package timego

import "time"

// Now gets time now
func Now() time.Time {
	return time.Now()
}

// NowUTC gets time now in UTC
func NowUTC() time.Time {
	return Now().UTC()
}
