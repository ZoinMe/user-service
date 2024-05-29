package service

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := us.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	return users, nil
}

func (us *UserService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := us.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	return user, nil
}

func (us *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := us.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return createdUser, nil
}

func (us *UserService) UpdateUser(ctx context.Context, updatedUser *model.User) (*model.User, error) {
	updatedUser, err := us.userRepository.UpdateUser(ctx, updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedUser, nil
}

func (us *UserService) DeleteUser(ctx context.Context, id uint) error {
	err := us.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}
