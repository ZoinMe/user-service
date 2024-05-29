package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type EducationService struct {
	educationRepository *repository.EducationRepository
}

func NewEducationService(educationRepository *repository.EducationRepository) *EducationService {
	return &EducationService{educationRepository}
}

func (es *EducationService) GetAllEducations(ctx context.Context) ([]*model.Education, error) {
	educations, err := es.educationRepository.GetAllEducations(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all educations: %v", err)
	}

	return educations, nil
}

func (es *EducationService) GetEducationByID(ctx context.Context, id uint) (*model.Education, error) {
	education, err := es.educationRepository.GetEducationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get education by ID: %v", err)
	}

	return education, nil
}

func (es *EducationService) CreateEducation(ctx context.Context, education *model.Education) (*model.Education, error) {
	createdEducation, err := es.educationRepository.CreateEducation(ctx, education)
	if err != nil {
		return nil, fmt.Errorf("failed to create education: %v", err)
	}

	return createdEducation, nil
}

func (es *EducationService) UpdateEducation(ctx context.Context, updatedEducation *model.Education) (*model.Education, error) {
	updatedEducation, err := es.educationRepository.UpdateEducation(ctx, updatedEducation)
	if err != nil {
		return nil, fmt.Errorf("failed to update education: %v", err)
	}

	return updatedEducation, nil
}

func (es *EducationService) DeleteEducation(ctx context.Context, id uint) error {
	err := es.educationRepository.DeleteEducation(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete education: %v", err)
	}

	return nil
}

func (es *EducationService) GetEducationsByUserID(ctx context.Context, userID uint) ([]*model.Education, error) {
	educations, err := es.educationRepository.GetEducationsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get educations by user ID: %v", err)
	}

	return educations, nil
}
