package util

import "github.com/satori/go.uuid"

func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
