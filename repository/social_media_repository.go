package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type SocialMediaRepository struct {
	DB *sql.DB
}

func NewSocialMediaRepository(db *sql.DB) *SocialMediaRepository {
	return &SocialMediaRepository{DB: db}
}

func (sr *SocialMediaRepository) GetAllSocialMedia(ctx context.Context) ([]*model.SocialMedia, error) {
	query := "SELECT id, name FROM social_media"
	rows, err := sr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all social media: %v", err)
	}
	defer rows.Close()

	var socialMedia []*model.SocialMedia
	for rows.Next() {
		var sm model.SocialMedia
		if err := rows.Scan(
			&sm.ID,
			&sm.Name,
		); err != nil {
			return nil, fmt.Errorf("failed to scan social media row: %v", err)
		}
		socialMedia = append(socialMedia, &sm)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading social media rows: %v", err)
	}

	return socialMedia, nil
}

func (sr *SocialMediaRepository) GetSocialMediaByID(ctx context.Context, id uint) (*model.SocialMedia, error) {
	query := "SELECT id, name FROM social_media WHERE id = ?"
	row := sr.DB.QueryRowContext(ctx, query, id)

	var sm model.SocialMedia
	if err := row.Scan(
		&sm.ID,
		&sm.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("social media not found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan social media row: %v", err)
	}

	return &sm, nil
}

func (sr *SocialMediaRepository) CreateSocialMedia(ctx context.Context, sm *model.SocialMedia) (*model.SocialMedia, error) {
	query := "INSERT INTO social_media (name) VALUES (?)"
	result, err := sr.DB.ExecContext(ctx, query, sm.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create social media: %v", err)
	}
	sm.ID, _ = result.LastInsertId()
	return sm, nil
}

func (sr *SocialMediaRepository) UpdateSocialMedia(ctx context.Context, updatedSM *model.SocialMedia) (*model.SocialMedia, error) {
	query := "UPDATE social_media SET name = ? WHERE id = ?"
	_, err := sr.DB.ExecContext(ctx, query, updatedSM.Name, updatedSM.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update social media: %v", err)
	}
	return updatedSM, nil
}

func (sr *SocialMediaRepository) DeleteSocialMedia(ctx context.Context, id uint) error {
	query := "DELETE FROM social_media WHERE id = ?"
	_, err := sr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete social media: %v", err)
	}
	return nil
}
