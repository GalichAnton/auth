package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GalichAnton/auth/internal/repository"
	userService "github.com/GalichAnton/auth/internal/services/user"
	"github.com/GalichAnton/platform_common/pkg/db"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *TestSuite) TestDelete() {
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

		id      = gofakeit.Int64()
		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name               string
		args               args
		want               error
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
			want: nil,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.DeleteMock.Expect(ctx, id).Return(nil)
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
			want: nil,
			err:  repoErr,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.DeleteMock.Expect(ctx, id).Return(repoErr)
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

				err := service.Delete(tt.args.ctx, tt.args.id)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
				}
			},
		)
	}
}
