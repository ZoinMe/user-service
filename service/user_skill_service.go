package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type UserSkillService struct {
	userSkillRepository *repository.UserSkillRepository
}

func NewUserSkillService(userSkillRepository *repository.UserSkillRepository) *UserSkillService {
	return &UserSkillService{userSkillRepository}
}

func (uss *UserSkillService) GetAllUserSkills(ctx context.Context) ([]*model.UserSkill, error) {
	userSkills, err := uss.userSkillRepository.GetAllUserSkills(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user skills: %v", err)
	}
	return userSkills, nil
}

func (uss *UserSkillService) GetUserSkillByID(ctx context.Context, id uint) (*model.UserSkill, error) {
	userSkill, err := uss.userSkillRepository.GetUserSkillByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user skill by ID: %v", err)
	}
	return userSkill, nil
}

func (uss *UserSkillService) CreateUserSkill(ctx context.Context, userSkill *model.UserSkill) (*model.UserSkill, error) {
	createdUserSkill, err := uss.userSkillRepository.CreateUserSkill(ctx, userSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to create user skill: %v", err)
	}
	return createdUserSkill, nil
}

func (uss *UserSkillService) UpdateUserSkill(ctx context.Context, updatedUserSkill *model.UserSkill) (*model.UserSkill, error) {
	updatedUserSkill, err := uss.userSkillRepository.UpdateUserSkill(ctx, updatedUserSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to update user skill: %v", err)
	}
	return updatedUserSkill, nil
}

func (uss *UserSkillService) DeleteUserSkill(ctx context.Context, id uint) error {
	err := uss.userSkillRepository.DeleteUserSkill(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user skill: %v", err)
	}
	return nil
}
