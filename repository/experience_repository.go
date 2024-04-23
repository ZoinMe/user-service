package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type ExperienceRepository struct {
	DB *sql.DB
}

func NewExperienceRepository(db *sql.DB) *ExperienceRepository {
	return &ExperienceRepository{DB: db}
}

func (er *ExperienceRepository) GetAllExperiences(ctx context.Context) ([]*model.Experience, error) {
	query := "SELECT id, company_logo, designation, company, from_date, to_date, description, user_id FROM experiences"
	rows, err := er.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all experiences: %v", err)
	}
	defer rows.Close()

	var experiences []*model.Experience
	for rows.Next() {
		var experience model.Experience
		if err := rows.Scan(
			&experience.ID,
			&experience.CompanyLogo,
			&experience.Designation,
			&experience.Company,
			&experience.FromDate,
			&experience.ToDate,
			&experience.Description,
			&experience.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan experience row: %v", err)
		}
		experiences = append(experiences, &experience)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading experience rows: %v", err)
	}

	return experiences, nil
}

func (er *ExperienceRepository) GetExperienceByID(ctx context.Context, id uint) (*model.Experience, error) {
	query := "SELECT id, company_logo, designation, company, from_date, to_date, description, user_id FROM experiences WHERE id = ?"
	row := er.DB.QueryRowContext(ctx, query, id)

	var experience model.Experience
	if err := row.Scan(
		&experience.ID,
		&experience.CompanyLogo,
		&experience.Designation,
		&experience.Company,
		&experience.FromDate,
		&experience.ToDate,
		&experience.Description,
		&experience.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("experience not found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan experience row: %v", err)
	}

	return &experience, nil
}

func (er *ExperienceRepository) CreateExperience(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	query := "INSERT INTO experiences (company_logo, designation, company, from_date, to_date, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := er.DB.ExecContext(ctx, query, experience.CompanyLogo, experience.Designation, experience.Company, experience.FromDate, experience.ToDate, experience.Description, experience.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create experience: %v", err)
	}
	experience.ID, _ = result.LastInsertId()
	return experience, nil
}

func (er *ExperienceRepository) UpdateExperience(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error) {
	query := "UPDATE experiences SET company_logo = ?, designation = ?, company = ?, from_date = ?, to_date = ?, description = ?, user_id = ? WHERE id = ?"
	_, err := er.DB.ExecContext(ctx, query, updatedExperience.CompanyLogo, updatedExperience.Designation, updatedExperience.Company, updatedExperience.FromDate, updatedExperience.ToDate, updatedExperience.Description, updatedExperience.UserID, updatedExperience.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update experience: %v", err)
	}
	return updatedExperience, nil
}

func (er *ExperienceRepository) DeleteExperience(ctx context.Context, id uint) error {
	query := "DELETE FROM experiences WHERE id = ?"
	_, err := er.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete experience: %v", err)
	}
	return nil
}
