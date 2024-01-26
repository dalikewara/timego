package timego

import "time"

// ToString converts time.Time to string format
func ToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// ToUTC converts time.Time location to UTC
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// ToLocation converts time.Time location to specified one (such as "Asia/Jakarta", "America/Los_Angeles", etc)
func ToLocation(t time.Time, location string) (time.Time, error) {
	zone, err := time.LoadLocation(location)
	if err != nil {
		return t, err
	}

	t = ToUTC(t).In(zone)

	return t, nil
}
