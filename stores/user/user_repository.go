package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZoinMe/user-service/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Get(ctx context.Context) ([]*model.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at, designation, bio, profile_image, location FROM users"

	rows, err := ur.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Designation,
			&user.Bio,
			&user.ProfileImage,
			&user.Location,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading user rows: %v", err)
	}

	return users, nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at, designation, bio, profile_image, location FROM users WHERE id = ?"
	row := ur.DB.QueryRowContext(ctx, query, id)

	var user model.User

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Designation,
		&user.Bio,
		&user.ProfileImage,
		&user.Location,
	); err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	return &user, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := "INSERT INTO users (name, email, password, created_at, updated_at, designation, bio, profile_image, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := ur.DB.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Designation, user.Bio, user.ProfileImage, user.Location)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userID, _ := result.LastInsertId()
	user.ID = userID

	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, updatedUser *model.User) (*model.User, error) {
	updatedUser.UpdatedAt = time.Now()

	query := "UPDATE users SET name=?, email=?, password=?, created_at=?, updated_at=?, designation=?, bio=?, profile_image=?, location=? WHERE id=?"

	_, err := ur.DB.ExecContext(ctx, query, updatedUser.Name, updatedUser.Email, updatedUser.Password, updatedUser.CreatedAt, updatedUser.UpdatedAt, updatedUser.Designation, updatedUser.Bio, updatedUser.ProfileImage, updatedUser.Location, updatedUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedUser, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := ur.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}
