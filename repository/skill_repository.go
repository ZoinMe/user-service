package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type SkillRepository struct {
	DB *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *SkillRepository {
	return &SkillRepository{DB: db}
}

func (sr *SkillRepository) GetAllSkills(ctx context.Context) ([]*model.Skill, error) {
	var skills []*model.Skill
	if err := sr.DB.Find(&skills).Error; err != nil {
		return nil, fmt.Errorf("failed to get all skills: %v", err)
	}
	return skills, nil
}

func (sr *SkillRepository) GetSkillByID(ctx context.Context, id uint) (*model.Skill, error) {
	var skill model.Skill
	if err := sr.DB.First(&skill, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get skill by ID: %v", err)
	}
	return &skill, nil
}

func (sr *SkillRepository) CreateSkill(ctx context.Context, skill *model.Skill) (*model.Skill, error) {
	if err := sr.DB.Create(&skill).Error; err != nil {
		return nil, fmt.Errorf("failed to create skill: %v", err)
	}
	return skill, nil
}

func (sr *SkillRepository) UpdateSkill(ctx context.Context, updatedSkill *model.Skill) (*model.Skill, error) {
	if err := sr.DB.Save(&updatedSkill).Error; err != nil {
		return nil, fmt.Errorf("failed to update skill: %v", err)
	}
	return updatedSkill, nil
}

func (sr *SkillRepository) DeleteSkill(ctx context.Context, id uint) error {
	if err := sr.DB.Delete(&model.Skill{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete skill: %v", err)
	}
	return nil
}
