package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type ExperienceRepository struct {
	DB *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{DB: db}
}

func (er *ExperienceRepository) GetAllExperiences(ctx context.Context) ([]*model.Experience, error) {
	var experiences []*model.Experience
	if err := er.DB.Find(&experiences).Error; err != nil {
		return nil, fmt.Errorf("failed to get all experiences: %v", err)
	}
	return experiences, nil
}

func (er *ExperienceRepository) GetExperienceByID(ctx context.Context, id uint) (*model.Experience, error) {
	var experience model.Experience
	if err := er.DB.First(&experience, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get experience by ID: %v", err)
	}
	return &experience, nil
}

func (er *ExperienceRepository) CreateExperience(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	if err := er.DB.Create(&experience).Error; err != nil {
		return nil, fmt.Errorf("failed to create experience: %v", err)
	}
	return experience, nil
}

func (er *ExperienceRepository) UpdateExperience(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error) {
	if err := er.DB.Save(&updatedExperience).Error; err != nil {
		return nil, fmt.Errorf("failed to update experience: %v", err)
	}
	return updatedExperience, nil
}

func (er *ExperienceRepository) DeleteExperience(ctx context.Context, id uint) error {
	if err := er.DB.Delete(&model.Experience{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete experience: %v", err)
	}
	return nil
}
