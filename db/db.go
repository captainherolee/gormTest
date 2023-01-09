package db

import (
	"gormTest/models"
)

type DB interface {
	GetUser(email string) (models.User, error)
}
