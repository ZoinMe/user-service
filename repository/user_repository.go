package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := ur.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, updatedUser *model.User) (*model.User, error) {
	updatedUser.UpdatedAt = time.Now()
	if err := ur.DB.Save(&updatedUser).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return updatedUser, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	if err := ur.DB.Delete(&model.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
