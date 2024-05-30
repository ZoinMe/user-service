package skill

import (
	"context"
	"fmt"
	"github.com/ZoinMe/user-service/stores/skill"

	"github.com/ZoinMe/user-service/model"
)

type skillService struct {
	skillRepository *skill.SkillRepository
}

func NewSkillService(skillRepository *skill.SkillRepository) *skillService {
	return &skillService{skillRepository}
}

func (ss *skillService) Get(ctx context.Context) ([]*model.Skill, error) {
	skills, err := ss.skillRepository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all skills: %v", err)
	}

	return skills, nil
}

func (ss *skillService) GetByID(ctx context.Context, id uint) (*model.Skill, error) {
	skill, err := ss.skillRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get skill by ID: %v", err)
	}

	return skill, nil
}

func (ss *skillService) Create(ctx context.Context, skill *model.Skill) (*model.Skill, error) {
	createdSkill, err := ss.skillRepository.Create(ctx, skill)
	if err != nil {
		return nil, fmt.Errorf("failed to create skill: %v", err)
	}

	return createdSkill, nil
}

func (ss *skillService) Update(ctx context.Context, updatedSkill *model.Skill) (*model.Skill, error) {
	updatedSkill, err := ss.skillRepository.Update(ctx, updatedSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to update skill: %v", err)
	}

	return updatedSkill, nil
}

func (ss *skillService) Delete(ctx context.Context, id uint) error {
	err := ss.skillRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete skill: %v", err)
	}

	return nil
}
