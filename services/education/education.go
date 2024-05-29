package education

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/services"
	"github.com/ZoinMe/user-service/stores"
)

type educationService struct {
	educationRepository stores.Education
}

func NewEducationService(educationRepository services.Education) services.Education {
	return &educationService{educationRepository}
}

func (es *educationService) Get(ctx context.Context) ([]*model.Education, error) {
	educations, err := es.educationRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all educations: %v", err)
	}

	return educations, nil
}

func (es *educationService) GetByID(ctx context.Context, id uint) (*model.Education, error) {
	education, err := es.educationRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get education by ID: %v", err)
	}

	return education, nil
}

func (es *educationService) Create(ctx context.Context, education *model.Education) (*model.Education, error) {
	createdEducation, err := es.educationRepository.Create(ctx, education)
	if err != nil {
		return nil, fmt.Errorf("failed to create education: %v", err)
	}

	return createdEducation, nil
}

func (es *educationService) Update(ctx context.Context, updatedEducation *model.Education) (*model.Education, error) {
	updatedEducation, err := es.educationRepository.Update(ctx, updatedEducation)
	if err != nil {
		return nil, fmt.Errorf("failed to update education: %v", err)
	}

	return updatedEducation, nil
}

func (es *educationService) Delete(ctx context.Context, id uint) error {
	err := es.educationRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete education: %v", err)
	}

	return nil
}

func (es *educationService) GetByUserID(ctx context.Context, userID uint) ([]*model.Education, error) {
	educations, err := es.educationRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get educations by user ID: %v", err)
	}

	return educations, nil
}
