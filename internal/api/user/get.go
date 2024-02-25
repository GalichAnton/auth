package user

import (
	"context"

	"github.com/GalichAnton/auth/internal/converter"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

// Get ...
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToGRPCUser(user),
	}, nil
}
