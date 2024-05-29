package user

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/stores/user"
)

type UserService struct {
	userRepository *user.UserRepository
}

func NewUserService(userRepository *user.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) Get(ctx context.Context) ([]*model.User, error) {
	users, err := us.userRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	return users, nil
}

func (us *UserService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := us.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	return user, nil
}

func (us *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := us.userRepository.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return createdUser, nil
}

func (us *UserService) Update(ctx context.Context, updatedUser *model.User) (*model.User, error) {
	updatedUser, err := us.userRepository.Update(ctx, updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedUser, nil
}

func (us *UserService) Delete(ctx context.Context, id uint) error {
	err := us.userRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}
