package timego

import "time"

func ReplaceYear(t time.Time, year int) time.Time {
	return time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceMonth(t time.Time, month int) time.Time {
	return time.Date(t.Year(), time.Month(month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceDay(t time.Time, day int) time.Time {
	return time.Date(t.Year(), t.Month(), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceMinute(t time.Time, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceSecond(t time.Time, second int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second, t.Nanosecond(), t.Location())
}

func ReplaceNanoSecond(t time.Time, nanoSecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), nanoSecond, t.Location())
}

func ReplaceDate(t time.Time, year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ReplaceTime(t time.Time, hour int, minute int, second int, nanoSecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, nanoSecond, t.Location())
}
