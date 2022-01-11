package allure

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

// GetUUID ...
func GetUUID() uuid.UUID {
	u, _ := uuid.NewUUID()
	return u
}

// GetMD5Hash ...
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// GetNow returns time.Now casted to int64 and time.Millisecond
func GetNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func matchInnerSteps(steps []*Step) []*Step {
	stepsTemp := make(map[string][]*Step)

	for idx := 0; idx < len(steps); idx++ {
		if steps[idx].Parent != "" {
			stepsTemp[steps[idx].Parent] = append(stepsTemp[steps[idx].Parent], steps[idx])
			steps = append(steps[:idx], steps[idx+1:]...)
			idx--
		}
	}

	for _, step := range stepsTemp {
		for idx := range step {
			step[idx].Steps = stepsTemp[step[idx].GetUUID()]
		}
	}

	for idx := 0; idx < len(steps); idx++ {
		steps[idx].Steps = stepsTemp[steps[idx].GetUUID()]
	}
	return steps
}
