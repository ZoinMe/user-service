package notification

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/services"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationService services.Notification
}

func NewNotificationHandler(notificationService services.Notification) *NotificationHandler {
	return &NotificationHandler{notificationService}
}

func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
	notifications, err := h.notificationService.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandler) GetNotificationByID(c *gin.Context) {
	notificationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	notification, err := h.notificationService.GetByID(c.Request.Context(), notificationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Notification with ID %d not found", notificationID)})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification model.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newNotification, err := h.notificationService.Create(c.Request.Context(), &notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newNotification)
}

func (h *NotificationHandler) UpdateNotification(c *gin.Context) {
	notificationID := c.Param("id")

	var updatedNotification model.Notification
	if err := c.ShouldBindJSON(&updatedNotification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedNotification.ID = notificationID

	notification, err := h.notificationService.Update(c.Request.Context(), &updatedNotification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	notificationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	err = h.notificationService.Delete(c.Request.Context(), notificationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}

func (h *NotificationHandler) GetNotificationsByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	notifications, err := h.notificationService.GetByUserID(c.Request.Context(), userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
