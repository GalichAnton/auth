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
	"github.com/GalichAnton/platform_common/pkg/db/pg"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *TestSuite) TestUpdate() {
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type txTransactorMockFunc func() db.Transactor

	type args struct {
		ctx  context.Context
		id   int64
		info *user.Info
	}

	var (
		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    TxMock
		ctx    = context.Background()

		id             = gofakeit.Int64()
		userRole int32 = 1
		name           = gofakeit.Name()
		email          = gofakeit.Email()
		password       = gofakeit.Password(true, true, true, true, true, 4)
		repoErr        = fmt.Errorf("repo error")
		txError        = errors.Wrap(repoErr, "failed executing code inside transaction")

		testInfo = &user.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     userRole,
		}

		logInfo = &log.Info{
			Action:   "update",
			EntityID: id,
		}
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
				ctx:  ctx,
				id:   id,
				info: testInfo,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.UpdateMock.Expect(pg.MakeContextTx(ctx, &txM), id, testInfo).Return(nil)
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
				ctx:  ctx,
				id:   id,
				info: testInfo,
			},
			want: nil,
			err:  txError,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.UpdateMock.Expect(pg.MakeContextTx(ctx, &txM), id, testInfo).Return(repoErr)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			txTransactorMock: func() db.Transactor {
				return s.txTransactorMock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(
			tt.name, func(t *testing.T) {
				userRepoMock := tt.userRepositoryMock()
				logRepoMock := tt.logRepositoryMock()
				txManagerMock := transaction.NewTransactionManager(tt.txTransactorMock())
				service := userService.NewService(userRepoMock, txManagerMock, logRepoMock)

				err := service.Update(tt.args.ctx, tt.args.id, tt.args.info)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
				}
			},
		)
	}
}
