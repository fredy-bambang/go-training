package repositories

import "context"

// User .
type User struct {
	Username string
	Password string
}

// UserRepository .
type UserRepository interface {
	Insert(ctx context.Context, user *User) error
}
