package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type EducationRepository struct {
	DB *gorm.DB
}

func NewEducationRepository(db *gorm.DB) *EducationRepository {
	return &EducationRepository{DB: db}
}

func (er *EducationRepository) GetAllEducations(ctx context.Context) ([]*model.Education, error) {
	var educations []*model.Education
	if err := er.DB.Find(&educations).Error; err != nil {
		return nil, fmt.Errorf("failed to get all educations: %v", err)
	}
	return educations, nil
}

func (er *EducationRepository) GetEducationByID(ctx context.Context, id uint) (*model.Education, error) {
	var education model.Education
	if err := er.DB.First(&education, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get education by ID: %v", err)
	}
	return &education, nil
}

func (er *EducationRepository) CreateEducation(ctx context.Context, education *model.Education) (*model.Education, error) {
	if err := er.DB.Create(&education).Error; err != nil {
		return nil, fmt.Errorf("failed to create education: %v", err)
	}
	return education, nil
}

func (er *EducationRepository) UpdateEducation(ctx context.Context, updatedEducation *model.Education) (*model.Education, error) {
	if err := er.DB.Save(&updatedEducation).Error; err != nil {
		return nil, fmt.Errorf("failed to update education: %v", err)
	}
	return updatedEducation, nil
}

func (er *EducationRepository) DeleteEducation(ctx context.Context, id uint) error {
	if err := er.DB.Delete(&model.Education{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete education: %v", err)
	}
	return nil
}
