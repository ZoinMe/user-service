package notification

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/user-service/model"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (nr *NotificationRepository) GetAll(ctx context.Context) ([]*model.Notification, error) {
	query := "SELECT id, user_id, message, type FROM notifications"

	rows, err := nr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all notifications: %v", err)
	}
	defer rows.Close()

	var notifications []*model.Notification

	for rows.Next() {
		var notification model.Notification

		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Message,
			&notification.Type,
		); err != nil {
			return nil, fmt.Errorf("failed to scan notification row: %v", err)
		}

		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading notification rows: %v", err)
	}

	return notifications, nil
}

func (nr *NotificationRepository) GetByID(ctx context.Context, id int) (*model.Notification, error) {
	query := "SELECT id, user_id, message, type FROM notifications WHERE id = ?"
	row := nr.DB.QueryRowContext(ctx, query, id)

	var notification model.Notification

	if err := row.Scan(
		&notification.ID,
		&notification.UserID,
		&notification.Message,
		&notification.Type,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("notification not found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan notification row: %v", err)
	}

	return &notification, nil
}

func (nr *NotificationRepository) Create(ctx context.Context, notification *model.Notification) (*model.Notification, error) {
	query := "INSERT INTO notifications (user_id, message, type) VALUES (?, ?, ?)"

	result, err := nr.DB.ExecContext(ctx, query, notification.UserID, notification.Message, notification.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification: %v", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve LastInsertId: %v", err)
	}

	return notification, nil
}

func (nr *NotificationRepository) Update(ctx context.Context, updatedNotification *model.Notification) (*model.Notification, error) {
	query := "UPDATE notifications SET user_id=?, message=?, type=? WHERE id=?"

	_, err := nr.DB.ExecContext(ctx, query, updatedNotification.UserID, updatedNotification.Message, updatedNotification.Type, updatedNotification.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update notification: %v", err)
	}

	return updatedNotification, nil
}

func (nr *NotificationRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM notifications WHERE id = ?"

	_, err := nr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %v", err)
	}

	return nil
}

func (nr *NotificationRepository) GetByUserID(ctx context.Context, userID string) ([]*model.Notification, error) {
	query := "SELECT id, user_id, message, type FROM notifications WHERE user_id = ?"

	rows, err := nr.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications by user ID: %v", err)
	}
	defer rows.Close()

	var notifications []*model.Notification

	for rows.Next() {
		var notification model.Notification

		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Message,
			&notification.Type,
		); err != nil {
			return nil, fmt.Errorf("failed to scan notification row: %v", err)
		}

		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading notification rows: %v", err)
	}

	return notifications, nil
}
