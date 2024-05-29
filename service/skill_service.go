package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
)

type SkillService struct {
	skillRepository *repository.SkillRepository
}

func NewSkillService(skillRepository *repository.SkillRepository) *SkillService {
	return &SkillService{skillRepository}
}

func (ss *SkillService) GetAllSkills(ctx context.Context) ([]*model.Skill, error) {
	skills, err := ss.skillRepository.GetAllSkills(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all skills: %v", err)
	}

	return skills, nil
}

func (ss *SkillService) GetSkillByID(ctx context.Context, id uint) (*model.Skill, error) {
	skill, err := ss.skillRepository.GetSkillByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get skill by ID: %v", err)
	}

	return skill, nil
}

func (ss *SkillService) CreateSkill(ctx context.Context, skill *model.Skill) (*model.Skill, error) {
	createdSkill, err := ss.skillRepository.CreateSkill(ctx, skill)
	if err != nil {
		return nil, fmt.Errorf("failed to create skill: %v", err)
	}

	return createdSkill, nil
}

func (ss *SkillService) UpdateSkill(ctx context.Context, updatedSkill *model.Skill) (*model.Skill, error) {
	updatedSkill, err := ss.skillRepository.UpdateSkill(ctx, updatedSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to update skill: %v", err)
	}

	return updatedSkill, nil
}

func (ss *SkillService) DeleteSkill(ctx context.Context, id uint) error {
	err := ss.skillRepository.DeleteSkill(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete skill: %v", err)
	}

	return nil
}
