package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type ExperienceService struct {
	experienceRepository *repository.ExperienceRepository
}

func NewExperienceService(experienceRepository *repository.ExperienceRepository) *ExperienceService {
	return &ExperienceService{experienceRepository}
}

func (es *ExperienceService) GetAllExperiences(ctx context.Context) ([]*model.Experience, error) {
	experiences, err := es.experienceRepository.GetAllExperiences(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all experiences: %v", err)
	}
	return experiences, nil
}

func (es *ExperienceService) GetExperienceByID(ctx context.Context, id uint) (*model.Experience, error) {
	experience, err := es.experienceRepository.GetExperienceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get experience by ID: %v", err)
	}
	return experience, nil
}

func (es *ExperienceService) CreateExperience(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	createdExperience, err := es.experienceRepository.CreateExperience(ctx, experience)
	if err != nil {
		return nil, fmt.Errorf("failed to create experience: %v", err)
	}
	return createdExperience, nil
}

func (es *ExperienceService) UpdateExperience(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error) {
	updatedExperience, err := es.experienceRepository.UpdateExperience(ctx, updatedExperience)
	if err != nil {
		return nil, fmt.Errorf("failed to update experience: %v", err)
	}
	return updatedExperience, nil
}

func (es *ExperienceService) DeleteExperience(ctx context.Context, id uint) error {
	err := es.experienceRepository.DeleteExperience(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete experience: %v", err)
	}
	return nil
}
