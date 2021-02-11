package handlers

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rickywinata/go-training/project/internal/decoders"
	"github.com/rickywinata/go-training/project/internal/endpoints"
	"github.com/rickywinata/go-training/project/internal/services"
)

// CreateUserHandler .
func CreateUserHandler(svc services.UserService) http.Handler {
	return httptransport.NewServer(
		endpoints.CreateUserEndpoint(svc),
		decoders.DecodeUserRequest,
		encodeResponse,
	)
}
