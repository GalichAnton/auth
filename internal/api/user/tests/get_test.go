package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	userApi "github.com/GalichAnton/auth/internal/api/user"
	modelService "github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/services"
	serviceMocks "github.com/GalichAnton/auth/internal/services/mocks"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id              = gofakeit.Int64()
		userRole  int32 = 1
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(true, true, true, true, true, 4)
		createdAt       = time.Now()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		modelUser = &modelService.User{
			ID: id,
			Info: modelService.Info{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     userRole,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{},
		}

		RPCInfo = &desc.UserInfo{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     desc.Role(userRole),
		}

		user = desc.User{
			Id:        id,
			Info:      RPCInfo,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: nil,
		}

		res = &desc.GetResponse{
			User: &user,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
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
				mock.GetMock.Expect(ctx, id).Return(modelUser, nil)
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
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
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

				response, err := api.Get(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, response)
			},
		)
	}
}
