package experience

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

func (er *ExperienceRepository) Get(ctx context.Context) ([]*model.Experience, error) {
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

func (er *ExperienceRepository) GetByID(ctx context.Context, id uint) (*model.Experience, error) {
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

func (er *ExperienceRepository) Create(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	query := "INSERT INTO experiences (company_logo, designation, company, from_date, to_date, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := er.DB.ExecContext(ctx, query, experience.CompanyLogo, experience.Designation, experience.Company, experience.FromDate, experience.ToDate, experience.Description, experience.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create experience: %v", err)
	}

	return experience, nil
}

func (er *ExperienceRepository) Update(ctx context.Context, updatedExperience *model.Experience) (*model.Experience, error) {
	query := "UPDATE experiences SET company_logo = ?, designation = ?, company = ?, from_date = ?, to_date = ?, description = ?, user_id = ? WHERE id = ?"

	_, err := er.DB.ExecContext(ctx, query, updatedExperience.CompanyLogo, updatedExperience.Designation, updatedExperience.Company, updatedExperience.FromDate, updatedExperience.ToDate, updatedExperience.Description, updatedExperience.UserID, updatedExperience.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update experience: %v", err)
	}

	return updatedExperience, nil
}

func (er *ExperienceRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM experiences WHERE id = ?"

	_, err := er.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete experience: %v", err)
	}

	return nil
}

func (er *ExperienceRepository) GetByUserID(ctx context.Context, userID string) ([]*model.Experience, error) {
	query := "SELECT id, company_logo, designation, company, from_date, to_date, description, user_id FROM experiences WHERE user_id = ?"

	rows, err := er.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get experiences by user ID: %v", err)
	}

	defer rows.Close()

	var experiences []*model.Experience

	for rows.Next() {
		var exp model.Experience

		if err := rows.Scan(
			&exp.ID,
			&exp.CompanyLogo,
			&exp.Designation,
			&exp.Company,
			&exp.FromDate,
			&exp.ToDate,
			&exp.Description,
			&exp.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan experience row: %v", err)
		}

		experiences = append(experiences, &exp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading experience rows: %v", err)
	}

	return experiences, nil
}
