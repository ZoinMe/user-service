package notification

import (
	"context"
	"fmt"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/stores/notification"
)

type NotificationService struct {
	notificationRepository *notification.NotificationRepository
}

func NewNotificationService(notificationRepository *notification.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository}
}

func (ns *NotificationService) Get(ctx context.Context) ([]*model.Notification, error) {
	notifications, err := ns.notificationRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all notifications: %v", err)
	}

	return notifications, nil
}

func (ns *NotificationService) GetByID(ctx context.Context, id int) (*model.Notification, error) {
	notification, err := ns.notificationRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification by ID: %v", err)
	}

	return notification, nil
}

func (ns *NotificationService) Create(ctx context.Context, notification *model.Notification) (*model.Notification, error) {
	createdNotification, err := ns.notificationRepository.Create(ctx, notification)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification: %v", err)
	}

	return createdNotification, nil
}

func (ns *NotificationService) Update(ctx context.Context, updatedNotification *model.Notification) (*model.Notification, error) {
	updatedNotification, err := ns.notificationRepository.Update(ctx, updatedNotification)
	if err != nil {
		return nil, fmt.Errorf("failed to update notification: %v", err)
	}

	return updatedNotification, nil
}

func (ns *NotificationService) Delete(ctx context.Context, id int) error {
	err := ns.notificationRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %v", err)
	}

	return nil
}

func (ns *NotificationService) GetByUserID(ctx context.Context, userID string) ([]*model.Notification, error) {
	notifications, err := ns.notificationRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications by user ID: %v", err)
	}

	return notifications, nil
}
