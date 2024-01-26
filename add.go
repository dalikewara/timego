package timego

import "time"

// AddYears adds the given number of years to time.Time
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// AddMonths adds the given number of months to time.Time
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddDays adds the given number of days to time.Time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddHours adds the given number of hours to time.Time
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}

// AddMinutes adds the given number of minutes to time.Time
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

// AddSeconds adds the given number of seconds to time.Time
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Second * time.Duration(seconds))
}

// AddNanoSeconds adds the given number of nanoseconds to time.Time
func AddNanoSeconds(t time.Time, nanoSeconds int) time.Time {
	return t.Add(time.Nanosecond * time.Duration(nanoSeconds))
}

// AddDuration adds the given time.Duration to time.Time
func AddDuration(t time.Time, duration time.Duration) time.Time {
	return t.Add(duration)
}
