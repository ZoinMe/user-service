package userSkill

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type UserSkillRepository struct {
	DB *sql.DB
}

func NewUserSkillRepository(db *sql.DB) *UserSkillRepository {
	return &UserSkillRepository{DB: db}
}

func (usr *UserSkillRepository) Get(ctx context.Context) ([]*model.UserSkill, error) {
	query := "SELECT id, user_id, skill_id FROM user_skills"

	rows, err := usr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user skills: %v", err)
	}

	defer rows.Close()

	var userSkills []*model.UserSkill

	for rows.Next() {
		var userSkill model.UserSkill

		if err := rows.Scan(
			&userSkill.ID,
			&userSkill.UserID,
			&userSkill.SkillID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user skill row: %v", err)
		}

		userSkills = append(userSkills, &userSkill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading user skill rows: %v", err)
	}

	return userSkills, nil
}

func (usr *UserSkillRepository) GetByID(ctx context.Context, id uint) (*model.UserSkill, error) {
	query := "SELECT id, user_id, skill_id FROM user_skills WHERE id = ?"
	row := usr.DB.QueryRowContext(ctx, query, id)

	var userSkill model.UserSkill

	if err := row.Scan(
		&userSkill.ID,
		&userSkill.UserID,
		&userSkill.SkillID,
	); err != nil {
		return nil, fmt.Errorf("failed to get user skill by ID: %v", err)
	}

	return &userSkill, nil
}

func (usr *UserSkillRepository) Create(ctx context.Context, userSkill *model.UserSkill) (*model.UserSkill, error) {
	query := "INSERT INTO user_skills (user_id, skill_id) VALUES (?, ?)"

	result, err := usr.DB.ExecContext(ctx, query, userSkill.UserID, userSkill.SkillID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user skill: %v", err)
	}

	userSkillID, _ := result.LastInsertId()
	userSkill.ID = userSkillID

	return userSkill, nil
}

func (usr *UserSkillRepository) Update(ctx context.Context, updatedUserSkill *model.UserSkill) (*model.UserSkill, error) {
	query := "UPDATE user_skills SET user_id=?, skill_id=? WHERE id=?"

	_, err := usr.DB.ExecContext(ctx, query, updatedUserSkill.UserID, updatedUserSkill.SkillID, updatedUserSkill.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update user skill: %v", err)
	}

	return updatedUserSkill, nil
}

func (usr *UserSkillRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM user_skills WHERE id = ?"

	_, err := usr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user skill: %v", err)
	}

	return nil
}

func (usr *UserSkillRepository) GetByUserID(ctx context.Context, userID uint) ([]*model.UserSkill, error) {
	query := "SELECT id, user_id, skill_id FROM user_skills WHERE user_id = ?"

	rows, err := usr.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user skills by user ID: %v", err)
	}

	defer rows.Close()

	var userSkills []*model.UserSkill

	for rows.Next() {
		var userSkill model.UserSkill

		if err := rows.Scan(&userSkill.ID, &userSkill.UserID, &userSkill.SkillID); err != nil {
			return nil, fmt.Errorf("failed to scan user skill row: %v", err)
		}

		userSkills = append(userSkills, &userSkill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading user skill rows: %v", err)
	}

	return userSkills, nil
}
