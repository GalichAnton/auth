package user

import (
	"context"

	"github.com/GalichAnton/auth/internal/converter"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update ...
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, req.GetId(), converter.ToServiceUserInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
