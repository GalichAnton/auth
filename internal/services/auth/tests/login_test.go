package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GalichAnton/auth/internal/config/env"
	"github.com/GalichAnton/auth/internal/models/claims"
	modelService "github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository"
	repoMocks "github.com/GalichAnton/auth/internal/repository/mocks"
	authService "github.com/GalichAnton/auth/internal/services/auth"
	"github.com/GalichAnton/auth/internal/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx       context.Context
		loginData *modelService.Login
	}

	var (
		ctx               = context.Background()
		mc                = minimock.NewController(t)
		email             = gofakeit.Email()
		name              = gofakeit.Name()
		role        int32 = 1
		pw                = gofakeit.Password(true, true, true, true, true, 4)
		hashedPW, _       = bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		tcm         env.TokensConfigMock
		repoErr     = fmt.Errorf("repo error")

		user = &modelService.User{
			Info: modelService.Info{
				Name:     name,
				Email:    email,
				Password: string(hashedPW),
				Role:     role,
			},
			CreatedAt: time.Now(),
		}

		loginData = &modelService.Login{
			Email:    email,
			Password: pw,
		}

		wrongLoginData = &modelService.Login{
			Email:    email,
			Password: "1234",
		}
	)

	refreshToken, _ := utils.GenerateToken(
		claims.UserClaims{
			Email: user.Info.Email,
			Role:  user.Info.Role,
		},
		[]byte(tcm.Config().RefreshSecret),
		tcm.Config().RefreshExpiration,
	)

	tests := []struct {
		name               string
		args               args
		want               *string
		err                error
		userRepositoryMock userRepositoryMockFunc
		tokensConfigMock   env.TokensConfigMock
	}{
		{
			name: "success case",
			args: args{
				ctx:       ctx,
				loginData: loginData,
			},
			want: &refreshToken,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetByEmailMock.Expect(ctx, email).Return(user, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx:       ctx,
				loginData: loginData,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetByEmailMock.Expect(ctx, email).Return(nil, repoErr)
				return mock
			},
		},
		{
			name: "wrong password case",
			args: args{
				ctx:       ctx,
				loginData: wrongLoginData,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetByEmailMock.Expect(ctx, email).Return(user, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				userRepositoryMock := tt.userRepositoryMock(mc)
				service := authService.NewService(userRepositoryMock, &tcm)

				response, err := service.Login(tt.args.ctx, tt.args.loginData)
				if err != nil {
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.Equal(t, tt.want, response)
			},
		)
	}
}
