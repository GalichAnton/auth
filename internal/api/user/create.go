package user

import (
	"context"
	"log"

	"github.com/GalichAnton/auth/internal/converter"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

// Create ...
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.userService.Create(ctx, converter.ToServiceUserToCreate(req.GetUser()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
