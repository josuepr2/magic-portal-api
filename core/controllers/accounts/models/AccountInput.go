package models

import uuid "github.com/satori/go.uuid"

type AccountInput struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Nickname  *string
	Email     string
	Password  string
	RoleID    uuid.UUID
}

type AccountUpdateInput struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Nickname  *string
	Email     string
}
