package experience

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/stores/experience"

	"github.com/ZoinMe/user-service/model"
)

type ExperienceService struct {
	experienceRepository *experience.ExperienceRepository
}

func NewExperienceService(experienceRepository *experience.ExperienceRepository) *ExperienceService {
	return &ExperienceService{experienceRepository}
}

func (es *ExperienceService) Get(ctx context.Context) ([]*model.Experience, error) {
	experiences, err := es.experienceRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all experiences: %v", err)
	}

	return experiences, nil
}

func (es *ExperienceService) GetByID(ctx context.Context, id uint) (*model.Experience, error) {
	experience, err := es.experienceRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get experience by ID: %v", err)
	}

	return experience, nil
}

func (es *ExperienceService) Create(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	createdExperience, err := es.experienceRepository.Create(ctx, experience)
	if err != nil {
		return nil, fmt.Errorf("failed to create experience: %v", err)
	}

	return createdExperience, nil
}

func (es *ExperienceService) Update(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error) {
	updatedExperience, err := es.experienceRepository.Update(ctx, updatedExperience)
	if err != nil {
		return nil, fmt.Errorf("failed to update experience: %v", err)
	}

	return updatedExperience, nil
}

func (es *ExperienceService) Delete(ctx context.Context, id uint) error {
	err := es.experienceRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete experience: %v", err)
	}

	return nil
}

func (es *ExperienceService) GetByUserID(ctx context.Context, userID string) ([]*model.Experience, error) {
	experiences, err := es.experienceRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get experiences by user ID: %v", err)
	}

	return experiences, nil
}
