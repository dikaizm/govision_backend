package helpers

import "github.com/google/uuid"

func GenerateUserID() string {
	// generate uuid
	newUUID := uuid.New()
	return newUUID.String()
}
