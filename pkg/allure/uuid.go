package allure

import "github.com/google/uuid"

// getUUID ...
func getUUID() uuid.UUID {
	u, _ := uuid.NewUUID()
	return u
}
