package inmemory

import (
	"context"

	"github.com/rickywinata/go-training/project/internal/repositories"
)

// InmemUserRepository implements InmemUserRepository.
type InmemUserRepository struct {
	Data []*repositories.User
}

// Insert inserts a new user.
func (r *InmemUserRepository) Insert(ctx context.Context, user *repositories.User) error {
	r.Data = append(r.Data, user)
	return nil
}
