package timego

import "time"

// Diff gets the different time between old time and current time in time.Duration
func Diff(oldTime time.Time, currentTime time.Time) time.Duration {
	ot, _ := FromString(ToString(oldTime, LayoutCustom26), LayoutCustom26)
	ct, _ := FromString(ToString(currentTime, LayoutCustom26), LayoutCustom26)
	return ct.Sub(ot)
}

// DiffInDays gets the different time between old time and current time in days
func DiffInDays(oldTime time.Time, currentTime time.Time) int {
	return int(Diff(oldTime, currentTime).Hours()) / 24
}

// DiffInSeconds gets the different time between old time and current time in seconds
func DiffInSeconds(oldTime time.Time, currentTime time.Time) float64 {
	return Diff(oldTime, currentTime).Seconds()
}
