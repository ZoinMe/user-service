package services

import (
	"context"

	"github.com/ZoinMe/user-service/model"
)

type Education interface {
	Get(ctx context.Context) ([]*model.Education, error)
	GetByID(ctx context.Context, id uint) (*model.Education, error)
	Create(ctx context.Context, education *model.Education) (*model.Education, error)
	Update(ctx context.Context, updatedEducation *model.Education) (*model.Education, error)
	Delete(ctx context.Context, id uint) error
	GetByUserID(ctx context.Context, userID string) ([]*model.Education, error)
}

type Experience interface {
	Get(ctx context.Context) ([]*model.Experience, error)
	GetByID(ctx context.Context, id uint) (*model.Experience, error)
	Create(ctx context.Context, experience *model.Experience) (*model.Experience, error)
	Update(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error)
	Delete(ctx context.Context, id uint) error
	GetByUserID(ctx context.Context, userID string) ([]*model.Experience, error)
}

type Skill interface {
	Get(ctx context.Context) ([]*model.Skill, error)
	GetByID(ctx context.Context, id uint) (*model.Skill, error)
	Create(ctx context.Context, skill *model.Skill) (*model.Skill, error)
	Update(ctx context.Context, updatedSkill *model.Skill) (*model.Skill, error)
	Delete(ctx context.Context, id uint) error
}

type User interface {
	Get(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, updatedUser *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
}

type UserSkill interface {
	Get(ctx context.Context) ([]*model.UserSkill, error)
	GetByID(ctx context.Context, id uint) (*model.UserSkill, error)
	Create(ctx context.Context, userSkill *model.UserSkill) (*model.UserSkill, error)
	Update(ctx context.Context, updatedUserSkill *model.UserSkill) (*model.UserSkill, error)
	Delete(ctx context.Context, id uint) error
	GetByUserID(ctx context.Context, userID string) ([]*model.UserSkill, error)
}

type Notification interface {
	Get(ctx context.Context) ([]*model.Notification, error)
	GetByID(ctx context.Context, id int) (*model.Notification, error)
	Create(ctx context.Context, notification *model.Notification) (*model.Notification, error)
	Update(ctx context.Context, updatedNotification *model.Notification) (*model.Notification, error)
	Delete(ctx context.Context, id int) error
	GetByUserID(ctx context.Context, userID string) ([]*model.Notification, error)
}
