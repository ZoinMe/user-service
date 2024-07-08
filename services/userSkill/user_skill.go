package userSkill

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/stores/userSkill"

	"github.com/ZoinMe/user-service/model"
)

type userSkillService struct {
	userSkillRepository *userSkill.UserSkillRepository
}

func NewUserSkillService(userSkillRepository *userSkill.UserSkillRepository) *userSkillService {
	return &userSkillService{userSkillRepository}
}

func (uss *userSkillService) Get(ctx context.Context) ([]*model.UserSkill, error) {
	userSkills, err := uss.userSkillRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user skills: %v", err)
	}

	return userSkills, nil
}

func (uss *userSkillService) GetByID(ctx context.Context, id uint) (*model.UserSkill, error) {
	userSkill, err := uss.userSkillRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user skill by ID: %v", err)
	}

	return userSkill, nil
}

func (uss *userSkillService) Create(ctx context.Context, userSkill *model.UserSkill) (*model.UserSkill, error) {
	createdUserSkill, err := uss.userSkillRepository.Create(ctx, userSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to create user skill: %v", err)
	}

	return createdUserSkill, nil
}

func (uss *userSkillService) Update(ctx context.Context, updatedUserSkill *model.UserSkill) (*model.UserSkill, error) {
	updatedUserSkill, err := uss.userSkillRepository.Update(ctx, updatedUserSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to update user skill: %v", err)
	}

	return updatedUserSkill, nil
}

func (uss *userSkillService) Delete(ctx context.Context, id uint) error {
	err := uss.userSkillRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user skill: %v", err)
	}

	return nil
}

func (uss *userSkillService) GetByUserID(ctx context.Context, userID string) ([]*model.UserSkill, error) {
	userSkills, err := uss.userSkillRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user skills by user ID: %v", err)
	}

	return userSkills, nil
}
