package converter

import (
	modelService "github.com/GalichAnton/auth/internal/models/user"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToGRPCUser ...
func ToGRPCUser(user *modelService.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToGRPCUserInfo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToGRPCUserInfo ...
func ToGRPCUserInfo(info modelService.Info) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     desc.Role(info.Role),
	}
}

// ToServiceUserInfo ...
func ToServiceUserInfo(info *desc.UserInfo) *modelService.Info {
	return &modelService.Info{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     int32(info.Role),
	}
}

// ToServiceUserToCreate ...
func ToServiceUserToCreate(info *desc.UserToCreate) *modelService.ToCreate {
	return &modelService.ToCreate{
		Info: modelService.Info{Name: info.Name,
			Email:    info.Email,
			Password: info.Password,
			Role:     int32(info.Role)},
		PasswordConfirm: info.PasswordConfirm,
	}
}
