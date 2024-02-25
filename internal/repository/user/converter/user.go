package converter

import (
	modelService "github.com/GalichAnton/auth/internal/models/user"
	modelRepo "github.com/GalichAnton/auth/internal/repository/user/model"
)

// ToServiceUser ...
func ToServiceUser(user *modelRepo.User) *modelService.User {
	return &modelService.User{
		ID:        user.ID,
		Info:      ToServiceUserInfo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToServiceUserInfo ...
func ToServiceUserInfo(info modelRepo.Info) modelService.Info {
	return modelService.Info{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role,
	}
}
