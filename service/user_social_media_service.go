package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type UserSocialMediaService struct {
	userSocialMediaRepository *repository.UserSocialMediaRepository
}

func NewUserSocialMediaService(userSocialMediaRepository *repository.UserSocialMediaRepository) *UserSocialMediaService {
	return &UserSocialMediaService{userSocialMediaRepository}
}

func (usms *UserSocialMediaService) GetAllUserSocialMedia(ctx context.Context) ([]*model.UserSocialMedia, error) {
	userSocialMedia, err := usms.userSocialMediaRepository.GetAllUserSocialMedia(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user social media: %v", err)
	}
	return userSocialMedia, nil
}

func (usms *UserSocialMediaService) GetUserSocialMediaByID(ctx context.Context, id uint) (*model.UserSocialMedia, error) {
	userSocialMedia, err := usms.userSocialMediaRepository.GetUserSocialMediaByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user social media by ID: %v", err)
	}
	return userSocialMedia, nil
}

func (usms *UserSocialMediaService) CreateUserSocialMedia(ctx context.Context, userSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	createdUserSocialMedia, err := usms.userSocialMediaRepository.CreateUserSocialMedia(ctx, userSocialMedia)
	if err != nil {
		return nil, fmt.Errorf("failed to create user social media: %v", err)
	}
	return createdUserSocialMedia, nil
}

func (usms *UserSocialMediaService) UpdateUserSocialMedia(ctx context.Context, updatedUserSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	updatedUserSocialMedia, err := usms.userSocialMediaRepository.UpdateUserSocialMedia(ctx, updatedUserSocialMedia)
	if err != nil {
		return nil, fmt.Errorf("failed to update user social media: %v", err)
	}
	return updatedUserSocialMedia, nil
}

func (usms *UserSocialMediaService) DeleteUserSocialMedia(ctx context.Context, id uint) error {
	err := usms.userSocialMediaRepository.DeleteUserSocialMedia(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user social media: %v", err)
	}
	return nil
}
