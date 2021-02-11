package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rickywinata/go-training/project/internal/services"
)

// CreateUserEndpoint .
func CreateUserEndpoint(svc services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cmd := request.(*services.CreateUserCommand)
		product, err := svc.CreateUser(ctx, cmd)
		return product, err
	}
}
