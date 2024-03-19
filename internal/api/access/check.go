package access

import (
	"context"

	desc "github.com/GalichAnton/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Check ...
func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
