package db

import "github.com/google/uuid"

func GetById(id string) string {
	return uuid.New().String()
}
