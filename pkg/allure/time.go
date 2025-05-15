package allure

import "time"

// GetNow returns [time.Now] as UNIX milliseconds
func GetNow() int64 {
	return time.Now().UnixMilli()
}
