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

// NewUserService initializes a new UserService with the provided UserRepository.
func NewUserService(userRepository *user.UserRepository) *UserService {
	return &UserService{userRepository}
}

// Get retrieves all users from the repository.
func (us *UserService) Get(ctx context.Context) ([]*model.User, error) {
	users, err := us.userRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	return users, nil
}

// GetByID retrieves a user by their ID from the repository.
func (us *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := us.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	return user, nil
}

// Create adds a new user to the repository.
func (us *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := us.userRepository.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return createdUser, nil
}

// Update modifies an existing user in the repository.
func (us *UserService) Update(ctx context.Context, updatedUser *model.User) (*model.User, error) {

	// Check if the user exists
	existingUser, err := us.userRepository.GetByID(ctx, updatedUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user for update: %v", err)
	}
	if existingUser == nil {
		return nil, fmt.Errorf("user with ID %d not found", updatedUser.ID)
	}

	// Perform the update
	updatedUser, err = us.userRepository.Update(ctx, updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedUser, nil
}

// Delete removes a user from the repository by their ID.
func (us *UserService) Delete(ctx context.Context, id uint) error {
	err := us.userRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}
