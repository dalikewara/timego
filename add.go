package timego

import "time"

func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}

func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Second * time.Duration(seconds))
}

func AddNanoSeconds(t time.Time, nanoSeconds int) time.Time {
	return t.Add(time.Nanosecond * time.Duration(nanoSeconds))
}

func AddDuration(t time.Time, duration time.Duration) time.Time {
	return t.Add(duration)
}
