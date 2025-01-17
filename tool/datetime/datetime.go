package datetime

import "time"

var TimeFormat = time.RFC3339

func Now() time.Time {
	return time.Now().Format(TimeFormat)
}
