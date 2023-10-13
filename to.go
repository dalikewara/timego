package timego

import "time"

func ToString(t time.Time, layout string) string {
	return t.Format(layout)
}
