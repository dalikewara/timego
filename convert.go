package timego

import "time"

func ConvertToUTC(t time.Time) time.Time {
	return t.UTC()
}

func ConvertToTimezone(t time.Time, timezone string) (time.Time, error) {
	zone, err := time.LoadLocation(timezone)
	if err != nil {
		return t, err
	}

	t = ConvertToUTC(t).In(zone)

	return t, nil
}
