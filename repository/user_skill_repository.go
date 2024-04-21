package repository

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/jinzhu/gorm"
)

type UserSkillRepository struct {
	DB *gorm.DB
}

func NewUserSkillRepository(db *gorm.DB) *UserSkillRepository {
	return &UserSkillRepository{DB: db}
}

func (usr *UserSkillRepository) GetAllUserSkills(ctx context.Context) ([]*model.UserSkill, error) {
	var userSkills []*model.UserSkill
	if err := usr.DB.Find(&userSkills).Error; err != nil {
		return nil, fmt.Errorf("failed to get all user skills: %v", err)
	}
	return userSkills, nil
}

func (usr *UserSkillRepository) GetUserSkillByID(ctx context.Context, id uint) (*model.UserSkill, error) {
	var userSkill model.UserSkill
	if err := usr.DB.First(&userSkill, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user skill by ID: %v", err)
	}
	return &userSkill, nil
}

func (usr *UserSkillRepository) CreateUserSkill(ctx context.Context, userSkill *model.UserSkill) (*model.UserSkill, error) {
	if err := usr.DB.Create(&userSkill).Error; err != nil {
		return nil, fmt.Errorf("failed to create user skill: %v", err)
	}
	return userSkill, nil
}

func (usr *UserSkillRepository) UpdateUserSkill(ctx context.Context, updatedUserSkill *model.UserSkill) (*model.UserSkill, error) {
	if err := usr.DB.Save(&updatedUserSkill).Error; err != nil {
		return nil, fmt.Errorf("failed to update user skill: %v", err)
	}
	return updatedUserSkill, nil
}

func (usr *UserSkillRepository) DeleteUserSkill(ctx context.Context, id uint) error {
	if err := usr.DB.Delete(&model.UserSkill{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user skill: %v", err)
	}
	return nil
}
