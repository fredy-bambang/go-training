package services

import (
	"context"

	"github.com/rickywinata/go-training/project/internal/repositories"
)

type (
	// CreateUserCommand represents the parameters for creating a product.
	CreateUserCommand struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// UserService is an interface for product operations.
type UserService interface {
	CreateUser(ctx context.Context, cmd *CreateUserCommand) (*repositories.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new product service.
func NewUserService(userRepo1 repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo1,
	}
}

func (s *userService) CreateUser(ctx context.Context, cmd *CreateUserCommand) (*repositories.User, error) {
	user := &repositories.User{
		Username: cmd.Username,
		Password: cmd.Password,
	}

	if err := s.userRepo.Insert(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
