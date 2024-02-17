package server

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserServer - .
type UserServer struct {
	desc.UnimplementedUserV1Server
	repository repository.UserRepository
}

// NewUserServer - .
func NewUserServer(repository repository.UserRepository) *UserServer {
	return &UserServer{
		repository: repository,
	}
}

// Create - .
func (s *UserServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	newUser := user.Info{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}

	id, err := s.repository.Create(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get - .
func (s *UserServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	dbUser, err := s.repository.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Id:        dbUser.ID,
		Name:      dbUser.Info.Name,
		Email:     dbUser.Info.Email,
		Role:      dbUser.Info.Role,
		CreatedAt: timestamppb.New(dbUser.CreatedAt),
		UpdatedAt: timestamppb.New(dbUser.UpdatedAt.Time),
	}, nil
}

// Update - .
func (s *UserServer) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	var userInfo user.Info

	if req.Name != nil {
		userInfo.Name = req.GetName().GetValue()
	}

	if req.Email != nil {
		userInfo.Email = req.GetEmail().GetValue()
	}

	if req.Role != desc.Role_UNKNOWN {
		userInfo.Role = req.GetRole()
	}

	err := s.repository.Update(ctx, req.GetId(), &userInfo)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Delete - .
func (s *UserServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := s.repository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
