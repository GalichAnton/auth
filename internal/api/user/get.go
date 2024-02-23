package user

import (
	"context"
	"log"

	"github.com/GalichAnton/auth/internal/converter"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

// Get ...
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n", user.ID, user.Info.Name,
		user.Info.Email, desc.Role(user.Info.Role), user.CreatedAt, user.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToGRPCUser(user),
	}, nil
}
