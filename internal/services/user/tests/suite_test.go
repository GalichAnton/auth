package tests

import (
	"testing"

	repoMocks "github.com/GalichAnton/auth/internal/repository/mocks"
	txMocks "github.com/GalichAnton/platform_common/pkg/db/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	userRepositoryMock *repoMocks.UserRepositoryMock
	logRepositoryMock  *repoMocks.LogRepositoryMock
	txTransactorMock   *txMocks.TransactorMock
	mc                 *minimock.Controller
}

func (s *TestSuite) SetupTest() {
	mc := minimock.NewController(s.T())
	s.mc = mc
	s.userRepositoryMock = repoMocks.NewUserRepositoryMock(mc)
	s.logRepositoryMock = repoMocks.NewLogRepositoryMock(mc)
	s.txTransactorMock = txMocks.NewTransactorMock(mc)
}

func TestApp(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
