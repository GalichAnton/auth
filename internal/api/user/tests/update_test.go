package tests

import (
	"context"
	"fmt"
	"testing"

	userApi "github.com/GalichAnton/auth/internal/api/user"
	modelService "github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/services"
	serviceMocks "github.com/GalichAnton/auth/internal/services/mocks"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id             = gofakeit.Int64()
		userRole int32 = 1
		name           = gofakeit.Name()
		email          = gofakeit.Email()
		password       = gofakeit.Password(true, true, true, true, true, 4)

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			Id: id,
			Info: &desc.UserInfo{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     desc.Role(userRole),
			},
		}

		info = &modelService.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     userRole,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, info).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, info).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				userServiceMock := tt.userServiceMock(mc)
				api := userApi.NewImplementation(userServiceMock)

				response, err := api.Update(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, response)
			},
		)
	}
}
