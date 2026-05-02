package service

import (
	"context"
	"vaultex/internal/model"
	"vaultex/internal/repository"
)

type Service interface {
	CreateUser(ctx context.Context, name, email string) (*model.User, error)
	FindUserByID(ctx context.Context, id string) (*model.User, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUserById(ctx context.Context, id string, name, email string) (*model.User, error)
	DeleteUserById(ctx context.Context, id string) error
}

type service struct {
	userRepo repository.UserRepository
}

func NewService(ur repository.UserRepository) Service {
	return &service{userRepo: ur}
}

func (s *service) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	err := s.userRepo.Create(ctx, user)

	return user, err
}

func (s *service) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *service) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

func (s *service) FindAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *service) UpdateUserById(ctx context.Context, id string, name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	updatedUser, err := s.userRepo.UpdateUserById(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *service) DeleteUserById(ctx context.Context, id string) error {
	return s.userRepo.DeleteUserById(ctx, id)
}
