package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type SocialMediaService struct {
	socialMediaRepository *repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository *repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{socialMediaRepository}
}

func (sms *SocialMediaService) GetAllSocialMedia(ctx context.Context) ([]*model.SocialMedia, error) {
	socialMedia, err := sms.socialMediaRepository.GetAllSocialMedia(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all social media: %v", err)
	}
	return socialMedia, nil
}

func (sms *SocialMediaService) GetSocialMediaByID(ctx context.Context, id uint) (*model.SocialMedia, error) {
	socialMedia, err := sms.socialMediaRepository.GetSocialMediaByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get social media by ID: %v", err)
	}
	return socialMedia, nil
}

func (sms *SocialMediaService) CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	createdSocialMedia, err := sms.socialMediaRepository.CreateSocialMedia(ctx, socialMedia)
	if err != nil {
		return nil, fmt.Errorf("failed to create social media: %v", err)
	}
	return createdSocialMedia, nil
}

func (sms *SocialMediaService) UpdateSocialMedia(ctx context.Context, updatedSocialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	updatedSocialMedia, err := sms.socialMediaRepository.UpdateSocialMedia(ctx, updatedSocialMedia)
	if err != nil {
		return nil, fmt.Errorf("failed to update social media: %v", err)
	}
	return updatedSocialMedia, nil
}

func (sms *SocialMediaService) DeleteSocialMedia(ctx context.Context, id uint) error {
	err := sms.socialMediaRepository.DeleteSocialMedia(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete social media: %v", err)
	}
	return nil
}
