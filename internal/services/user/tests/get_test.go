package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository"
	"github.com/GalichAnton/auth/internal/repository/user/model"
	userService "github.com/GalichAnton/auth/internal/services/user"
	"github.com/GalichAnton/platform_common/pkg/db"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *TestSuite) TestGet() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type txTransactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()

		id              = gofakeit.Int64()
		userRole  int32 = 1
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(true, true, true, true, true, 4)
		createdAt       = time.Now()
		repoErr         = fmt.Errorf("repo error")

		filter = model.Filter{
			ID: &id,
		}

		userInfo = user.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     userRole,
		}

		result = &user.User{
			ID:        id,
			Info:      userInfo,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *user.User
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: result,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.GetMock.Expect(ctx, filter).Return(result, nil)

			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			txTransactorMock: func() db.Transactor {
				return s.txTransactorMock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: result,
			err:  repoErr,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.GetMock.Expect(ctx, filter).Return(nil, repoErr)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			txTransactorMock: func() db.Transactor {
				return s.txTransactorMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				userRepoMock := tt.userRepositoryMock()
				logRepoMock := tt.logRepositoryMock()
				txManagerMock := transaction.NewTransactionManager(tt.txTransactorMock())
				service := userService.NewService(userRepoMock, txManagerMock, logRepoMock)

				response, err := service.Get(tt.args.ctx, tt.args.id)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
					return
				}
				s.Require().Equal(tt.want, response)
			},
		)
	}
}
