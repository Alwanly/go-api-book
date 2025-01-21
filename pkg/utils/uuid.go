package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return id.String()
}
