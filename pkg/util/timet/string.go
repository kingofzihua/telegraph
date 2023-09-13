package timet

import "time"

// ToString 转化为 string
func ToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.String()
}
