package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type UserSocialMediaRepository struct {
	DB *sql.DB
}

func NewUserSocialMediaRepository(db *sql.DB) *UserSocialMediaRepository {
	return &UserSocialMediaRepository{DB: db}
}

func (usr *UserSocialMediaRepository) GetAllUserSocialMedia(ctx context.Context) ([]*model.UserSocialMedia, error) {
	query := "SELECT id, user_id, url FROM user_social_media"
	rows, err := usr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user social media: %v", err)
	}
	defer rows.Close()

	var userSocialMedia []*model.UserSocialMedia
	for rows.Next() {
		var socialMedia model.UserSocialMedia
		if err := rows.Scan(
			&socialMedia.ID,
			&socialMedia.UserID,
			&socialMedia.URL,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user social media row: %v", err)
		}
		userSocialMedia = append(userSocialMedia, &socialMedia)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading user social media rows: %v", err)
	}

	return userSocialMedia, nil
}

func (usr *UserSocialMediaRepository) GetUserSocialMediaByID(ctx context.Context, id uint) (*model.UserSocialMedia, error) {
	query := "SELECT id, user_id, url FROM user_social_media WHERE id = ?"
	row := usr.DB.QueryRowContext(ctx, query, id)

	var socialMedia model.UserSocialMedia
	if err := row.Scan(
		&socialMedia.ID,
		&socialMedia.UserID,
		&socialMedia.URL,
	); err != nil {
		return nil, fmt.Errorf("failed to get user social media by ID: %v", err)
	}

	return &socialMedia, nil
}

func (usr *UserSocialMediaRepository) CreateUserSocialMedia(ctx context.Context, userSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	query := "INSERT INTO user_social_media (user_id, url) VALUES (?, ?)"
	result, err := usr.DB.ExecContext(ctx, query, userSocialMedia.UserID, userSocialMedia.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to create user social media: %v", err)
	}
	socialMediaID, _ := result.LastInsertId()
	userSocialMedia.ID = socialMediaID
	return userSocialMedia, nil
}

func (usr *UserSocialMediaRepository) UpdateUserSocialMedia(ctx context.Context, updatedUserSocialMedia *model.UserSocialMedia) (*model.UserSocialMedia, error) {
	query := "UPDATE user_social_media SET user_id=?, url=? WHERE id=?"
	_, err := usr.DB.ExecContext(ctx, query, updatedUserSocialMedia.UserID, updatedUserSocialMedia.URL, updatedUserSocialMedia.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update user social media: %v", err)
	}
	return updatedUserSocialMedia, nil
}

func (usr *UserSocialMediaRepository) DeleteUserSocialMedia(ctx context.Context, id uint) error {
	query := "DELETE FROM user_social_media WHERE id = ?"
	_, err := usr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user social media: %v", err)
	}
	return nil
}
