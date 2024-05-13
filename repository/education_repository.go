package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type EducationRepository struct {
	DB *sql.DB
}

func NewEducationRepository(db *sql.DB) *EducationRepository {
	return &EducationRepository{DB: db}
}

func (er *EducationRepository) GetAllEducations(ctx context.Context) ([]*model.Education, error) {
	query := "SELECT id, university_logo, university_name, degree, from_date, to_date, user_id FROM educations"
	rows, err := er.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all educations: %v", err)
	}
	defer rows.Close()

	var educations []*model.Education
	for rows.Next() {
		var education model.Education
		if err := rows.Scan(
			&education.ID,
			&education.UniversityLogo,
			&education.UniversityName,
			&education.Degree,
			&education.FromDate,
			&education.ToDate,
			&education.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan education row: %v", err)
		}
		educations = append(educations, &education)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading education rows: %v", err)
	}

	return educations, nil
}

func (er *EducationRepository) GetEducationByID(ctx context.Context, id uint) (*model.Education, error) {
	query := "SELECT id, university_logo, university_name, degree, from_date, to_date, user_id FROM educations WHERE id = ?"
	row := er.DB.QueryRowContext(ctx, query, id)

	var education model.Education
	if err := row.Scan(
		&education.ID,
		&education.UniversityLogo,
		&education.UniversityName,
		&education.Degree,
		&education.FromDate,
		&education.ToDate,
		&education.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("education not found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan education row: %v", err)
	}

	return &education, nil
}

func (er *EducationRepository) CreateEducation(ctx context.Context, education *model.Education) (*model.Education, error) {
	query := "INSERT INTO educations (university_logo, university_name, degree, from_date, to_date, user_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := er.DB.ExecContext(ctx, query, education.UniversityLogo, education.UniversityName, education.Degree, education.FromDate, education.ToDate, education.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create education: %v", err)
	}
	education.ID, _ = result.LastInsertId()
	return education, nil
}

func (er *EducationRepository) UpdateEducation(ctx context.Context, updatedEducation *model.Education) (*model.Education, error) {
	query := "UPDATE educations SET university_logo = ?, university_name = ?, degree = ?, from_date = ?, to_date = ?, user_id = ? WHERE id = ?"
	_, err := er.DB.ExecContext(ctx, query, updatedEducation.UniversityLogo, updatedEducation.UniversityName, updatedEducation.Degree, updatedEducation.FromDate, updatedEducation.ToDate, updatedEducation.UserID, updatedEducation.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update education: %v", err)
	}
	return updatedEducation, nil
}

func (er *EducationRepository) DeleteEducation(ctx context.Context, id uint) error {
	query := "DELETE FROM educations WHERE id = ?"
	_, err := er.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete education: %v", err)
	}
	return nil
}

func (er *EducationRepository) GetEducationsByUserID(ctx context.Context, userID uint) ([]*model.Education, error) {
	query := "SELECT id, university_logo, university_name, degree, from_date, to_date, user_id FROM educations WHERE user_id = ?"
	rows, err := er.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get educations by user ID: %v", err)
	}
	defer rows.Close()

	var educations []*model.Education
	for rows.Next() {
		var edu model.Education
		if err := rows.Scan(
			&edu.ID,
			&edu.UniversityLogo,
			&edu.UniversityName,
			&edu.Degree,
			&edu.FromDate,
			&edu.ToDate,
			&edu.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan education row: %v", err)
		}
		educations = append(educations, &edu)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading education rows: %v", err)
	}

	return educations, nil
}
