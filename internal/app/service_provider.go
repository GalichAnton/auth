package app

import (
	"context"
	"log"

	"github.com/GalichAnton/auth/internal/api/access"
	"github.com/GalichAnton/auth/internal/api/auth"
	"github.com/GalichAnton/auth/internal/api/user"
	"github.com/GalichAnton/auth/internal/config"
	"github.com/GalichAnton/auth/internal/config/env"
	"github.com/GalichAnton/auth/internal/repository"
	logRepository "github.com/GalichAnton/auth/internal/repository/log"
	roleRepository "github.com/GalichAnton/auth/internal/repository/role"
	userRepository "github.com/GalichAnton/auth/internal/repository/user"
	"github.com/GalichAnton/auth/internal/services"
	accessService "github.com/GalichAnton/auth/internal/services/access"
	authService "github.com/GalichAnton/auth/internal/services/auth"
	userService "github.com/GalichAnton/auth/internal/services/user"
	"github.com/GalichAnton/platform_common/pkg/closer"
	"github.com/GalichAnton/platform_common/pkg/db"
	"github.com/GalichAnton/platform_common/pkg/db/pg"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTPConfig
	swaggerConfig    config.SwaggerConfig
	tokensConfig     env.TokensConfig
	logConfig        env.LogConfig
	prometheusConfig env.PrometheusConfig

	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	roleRepository repository.RoleRepository

	userService   services.UserService
	authService   services.AuthService
	accessService services.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) TokensConfig() env.TokensConfig {
	if s.tokensConfig == nil {
		cfg, err := env.NewTokensConfig()
		if err != nil {
			log.Fatalf("failed to get tokens config: %s", err.Error())
		}

		s.tokensConfig = cfg
	}

	return s.tokensConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) LogConfig() config.LogConfig {
	if s.logConfig == nil {
		cfg, err := env.NewLogConfig()
		if err != nil {
			log.Fatalf("failed to get log config: %s", err.Error())
		}

		s.logConfig = cfg
	}

	return s.logConfig
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewLogRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

func (s *serviceProvider) RoleRepository(ctx context.Context) repository.RoleRepository {
	if s.roleRepository == nil {
		s.roleRepository = roleRepository.NewRoleRepository(s.DBClient(ctx))
	}

	return s.roleRepository
}

func (s *serviceProvider) UserService(ctx context.Context) services.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.LogRepository(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) services.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(ctx),
			s.TokensConfig(),
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) services.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.RoleRepository(ctx),
			s.TokensConfig(),
		)
	}

	return s.accessService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
