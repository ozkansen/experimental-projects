package repository

import (
	"context"
)

type User struct {
	ID       int
	Username string
	Passwd   string
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
}
