package tests

import (
	"context"
	"fmt"
	"testing"

	userApi "github.com/GalichAnton/auth/internal/api/user"
	"github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/services"
	serviceMocks "github.com/GalichAnton/auth/internal/services/mocks"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id                    = gofakeit.Int64()
		userRole        int32 = 1
		name                  = gofakeit.Name()
		email                 = gofakeit.Email()
		password              = gofakeit.Password(true, true, true, true, true, 4)
		passwordConfirm       = password

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			User: &desc.UserToCreate{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirm,
				Role:            desc.Role(userRole),
			},
		}

		info = user.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     userRole,
		}

		toCreate = &user.ToCreate{
			Info:            info,
			PasswordConfirm: passwordConfirm,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, toCreate).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, toCreate).Return(0, serviceErr)
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

				response, err := api.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, response)
			},
		)
	}
}
