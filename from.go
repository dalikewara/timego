package timego

import "time"

// FromString generates time.Time from string layout
func FromString(t string, layout string) (time.Time, error) {
	return time.Parse(layout, t)
}

// FromUnix generates time.Time from unix time
func FromUnix(unix int64) time.Time {
	return time.Unix(unix, 0).UTC()
}
