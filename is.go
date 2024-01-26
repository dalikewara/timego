package timego

import "time"

// IsExpired checks if the current time is expired or not
func IsExpired(currentTime time.Time, expiredTime time.Time) bool {
	if ToString(currentTime, LayoutCustom26) > ToString(expiredTime, LayoutCustom26) {
		return true
	}

	return false
}
