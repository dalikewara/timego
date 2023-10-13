package timego

import "time"

func Diff(currentTime time.Time, oldTime time.Time) time.Duration {
	ct, _ := FromString(ToString(currentTime, LayoutCustom26), LayoutCustom26)
	ot, _ := FromString(ToString(oldTime, LayoutCustom26), LayoutCustom26)
	return ct.Sub(ot)
}

func DiffInDays(currentTime time.Time, oldTime time.Time) int {
	return int(Diff(currentTime, oldTime).Hours()) / 24
}

func DiffTotalSeconds(currentTime time.Time, oldTime time.Time) float64 {
	return Diff(currentTime, oldTime).Seconds()
}
