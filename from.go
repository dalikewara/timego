package timego

import "time"

func FromString(t string, layout string) (time.Time, error) {
	return time.Parse(layout, t)
}

func FromUnix(unix int64) (time.Time, error) {
	return time.Unix(unix, 0), nil
}
