package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type BaseAccount struct {
	ID uuid.UUID
	FirstName string
	LastName string
	Nickname *string
	Email string
	RoleID uuid.UUID
}

type Account struct {
	ID uuid.UUID
	FirstName string
	LastName string
	Nickname *string
	Email string
	Password string
	CreatedAt time.Time
	DeletedAd time.Time
	RoleID uuid.UUID
}
