package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GalichAnton/auth/internal/models/log"
	"github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository"
	userService "github.com/GalichAnton/auth/internal/services/user"
	"github.com/GalichAnton/platform_common/pkg/db"
	txMocks "github.com/GalichAnton/platform_common/pkg/db/mocks"
	"github.com/GalichAnton/platform_common/pkg/db/pg"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *TestSuite) TestCreate() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type txTransactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		req *user.ToCreate
	}

	var (
		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    txMocks.TxMock
		ctx    = context.Background()

		id                    = gofakeit.Int64()
		userRole        int32 = 1
		name                  = gofakeit.Name()
		email                 = gofakeit.Email()
		password              = gofakeit.Password(true, true, true, true, true, 4)
		passwordConfirm       = password

		repoErr     = fmt.Errorf("repo error")
		passwordErr = fmt.Errorf("password and password confirmation do not match")
		txError     = errors.Wrap(repoErr, "failed executing code inside transaction")

		userInfo = &user.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     userRole,
		}

		userToCreate = &user.ToCreate{
			Info:            *userInfo,
			PasswordConfirm: passwordConfirm,
		}

		failedUserToCreate = &user.ToCreate{
			Info:            *userInfo,
			PasswordConfirm: gofakeit.Password(true, true, true, true, true, 4),
		}

		logInfo = &log.Info{
			Action:   "create",
			EntityID: id,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: userToCreate,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), userInfo).Return(id, nil)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), logInfo).Return(nil)
			},
			txTransactorMock: func() db.Transactor {
				return s.txTransactorMock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: userToCreate,
			},
			want: 0,
			err:  txError,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), userInfo).Return(0, repoErr)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			txTransactorMock: func() db.Transactor {
				return s.txTransactorMock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
			},
		},
		{
			name: "passwords not equal",
			args: args{
				ctx: ctx,
				req: failedUserToCreate,
			},
			want: 0,
			err:  passwordErr,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock
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

				response, err := service.Create(tt.args.ctx, tt.args.req)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
					return
				}
				s.Require().Equal(tt.want, response)
			},
		)
	}
}
