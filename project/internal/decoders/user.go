package decoders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rickywinata/go-training/project/internal/services"
)

// DecodeUserRequest .
func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var cmd services.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		return nil, err
	}
	return &cmd, nil
}
