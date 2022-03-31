package allure

import "time"

// GetNow returns time.Now casted to int64 and time.Millisecond
func GetNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
