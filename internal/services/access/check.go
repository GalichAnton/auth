package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/GalichAnton/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const authKey = "authorization"

func (s *service) Check(ctx context.Context, address string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md[authKey]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(s.tokensConfig.Config().AccessSecret))
	if err != nil {
		return errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[address]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

func (s *service) accessibleRoles(ctx context.Context) (map[string]int32, error) {
	var accessibleRoles map[string]int32

	if accessibleRoles == nil {
		accessibleRoles = make(map[string]int32)

		permissions, err := s.roleRepository.GetAllRolePermissions(ctx)
		if err != nil {
			return nil, err
		}

		for _, perm := range permissions {
			accessibleRoles[perm.Permission] = perm.RoleID
		}
	}

	return accessibleRoles, nil
}
