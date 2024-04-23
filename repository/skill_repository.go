package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type SkillRepository struct {
	DB *sql.DB
}

func NewSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{DB: db}
}

func (sr *SkillRepository) GetAllSkills(ctx context.Context) ([]*model.Skill, error) {
	query := "SELECT id, name, count FROM skills"
	rows, err := sr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all skills: %v", err)
	}
	defer rows.Close()

	var skills []*model.Skill
	for rows.Next() {
		var skill model.Skill
		if err := rows.Scan(
			&skill.ID,
			&skill.Name,
			&skill.Count,
		); err != nil {
			return nil, fmt.Errorf("failed to scan skill row: %v", err)
		}
		skills = append(skills, &skill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading skill rows: %v", err)
	}

	return skills, nil
}

func (sr *SkillRepository) GetSkillByID(ctx context.Context, id uint) (*model.Skill, error) {
	query := "SELECT id, name, count FROM skills WHERE id = ?"
	row := sr.DB.QueryRowContext(ctx, query, id)

	var skill model.Skill
	if err := row.Scan(
		&skill.ID,
		&skill.Name,
		&skill.Count,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("skill not found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan skill row: %v", err)
	}

	return &skill, nil
}

func (sr *SkillRepository) CreateSkill(ctx context.Context, skill *model.Skill) (*model.Skill, error) {
	query := "INSERT INTO skills (name, count) VALUES (?, ?)"
	result, err := sr.DB.ExecContext(ctx, query, skill.Name, skill.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to create skill: %v", err)
	}
	skill.ID, _ = result.LastInsertId()
	return skill, nil
}

func (sr *SkillRepository) UpdateSkill(ctx context.Context, updatedSkill *model.Skill) (*model.Skill, error) {
	query := "UPDATE skills SET name = ?, count = ? WHERE id = ?"
	_, err := sr.DB.ExecContext(ctx, query, updatedSkill.Name, updatedSkill.Count, updatedSkill.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update skill: %v", err)
	}
	return updatedSkill, nil
}

func (sr *SkillRepository) DeleteSkill(ctx context.Context, id uint) error {
	query := "DELETE FROM skills WHERE id = ?"
	_, err := sr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete skill: %v", err)
	}
	return nil
}
