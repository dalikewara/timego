package timego

import "time"

func Now() time.Time {
	return time.Now()
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func NowAddDays(days int) time.Time {
	return Now().AddDate(0, 0, days)
}

func NowUTCAddDays(days int) time.Time {
	return NowUTC().AddDate(0, 0, days)
}
