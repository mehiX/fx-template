package fxhttpserver

import "github.com/google/uuid"

func generateRequestId() string {
	return uuid.New().String()
}
