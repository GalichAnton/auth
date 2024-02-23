package user

import (
	"context"
	"log"

	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete ...
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	err := i.userService.Delete(ctx, id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("deleted user with id: %d", id)

	return &emptypb.Empty{}, nil
}
