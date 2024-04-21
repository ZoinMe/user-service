package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type UserSocialMediaRepository struct {
	DB *gorm.DB
}

func NewUserSocialMediaRepository(db *gorm.DB) *UserSocialMediaRepository {
	return &UserSocialMediaRepository{DB: db}
}

func (usr *UserSocialMediaRepository) GetAllUserSocialMedia(ctx context.Context) ([]*model.UserSocialMedia, error) {
	var userSocialMedia []*model.UserSocialMedia
	if err := usr.DB.Find(&userSocialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to get all user social media: %v", err)
	}
	return userSocialMedia, nil
}

func (usr *UserSocialMediaRepository) GetUserSocialMediaByID(ctx context.Context, id uint) (*model.UserSocialMedia, error) {
	var userSocialMedia model.UserSocialMedia
	if err := usr.DB.First(&userSocialMedia, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user social media by ID: %v", err)
	}
	return &userSocialMedia, nil
}

func (usr *UserSocialMediaRepository) CreateUserSocialMedia(ctx context.Context, userSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	if err := usr.DB.Create(&userSocialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to create user social media: %v", err)
	}
	return userSocialMedia, nil
}

func (usr *UserSocialMediaRepository) UpdateUserSocialMedia(ctx context.Context, updatedUserSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	if err := usr.DB.Save(&updatedUserSocialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to update user social media: %v", err)
	}
	return updatedUserSocialMedia, nil
}

func (usr *UserSocialMediaRepository) DeleteUserSocialMedia(ctx context.Context, id uint) error {
	if err := usr.DB.Delete(&model.UserSocialMedia{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user social media: %v", err)
	}
	return nil
}
