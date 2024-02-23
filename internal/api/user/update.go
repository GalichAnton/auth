package user

import (
	"context"
	"log"

	"github.com/GalichAnton/auth/internal/converter"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update ...
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	err := i.userService.Update(ctx, id, converter.ToServiceUserInfo(req.GetInfo()))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("updated user with id: %d", id)

	return &emptypb.Empty{}, nil
}
