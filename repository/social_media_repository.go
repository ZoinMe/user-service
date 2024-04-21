package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{DB: db}
}

func (sr *SocialMediaRepository) GetAllSocialMedia(ctx context.Context) ([]*model.SocialMedia, error) {
	var socialMedia []*model.SocialMedia
	if err := sr.DB.Find(&socialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to get all social media: %v", err)
	}
	return socialMedia, nil
}

func (sr *SocialMediaRepository) GetSocialMediaByID(ctx context.Context, id uint) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	if err := sr.DB.First(&socialMedia, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get social media by ID: %v", err)
	}
	return &socialMedia, nil
}

func (sr *SocialMediaRepository) CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := sr.DB.Create(&socialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to create social media: %v", err)
	}
	return socialMedia, nil
}

func (sr *SocialMediaRepository) UpdateSocialMedia(ctx context.Context, updatedSocialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := sr.DB.Save(&updatedSocialMedia).Error; err != nil {
		return nil, fmt.Errorf("failed to update social media: %v", err)
	}
	return updatedSocialMedia, nil
}

func (sr *SocialMediaRepository) DeleteSocialMedia(ctx context.Context, id uint) error {
	if err := sr.DB.Delete(&model.SocialMedia{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete social media: %v", err)
	}
	return nil
}
