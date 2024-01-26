package timego

import "time"

// ReplaceYear replaces time.Time year
func ReplaceYear(t time.Time, year int) time.Time {
	return time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceMonth replaces time.Time month
func ReplaceMonth(t time.Time, month int) time.Time {
	return time.Date(t.Year(), time.Month(month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceDay replaces time.Time day
func ReplaceDay(t time.Time, day int) time.Time {
	return time.Date(t.Year(), t.Month(), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceHour replaces time.Time hour
func ReplaceHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceMinute replaces time.Time minute
func ReplaceMinute(t time.Time, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceSecond replaces time.Time second
func ReplaceSecond(t time.Time, second int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second, t.Nanosecond(), t.Location())
}

// ReplaceNanoSecond replaces time.Time nanosecond
func ReplaceNanoSecond(t time.Time, nanoSecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), nanoSecond, t.Location())
}

// ReplaceDate replaces time.Time year, month and day
func ReplaceDate(t time.Time, year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceTime replaces time.Time hour, minute and nanosecond
func ReplaceTime(t time.Time, hour int, minute int, second int, nanoSecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, nanoSecond, t.Location())
}
