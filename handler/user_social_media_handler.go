package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
)

type UserSocialMediaHandler struct {
	userSocialMediaService *service.UserSocialMediaService
}

func NewUserSocialMediaHandler(userSocialMediaService *service.UserSocialMediaService) *UserSocialMediaHandler {
	return &UserSocialMediaHandler{userSocialMediaService}
}

func (h *UserSocialMediaHandler) GetAllUserSocialMedia(c *gin.Context) {
	userSocialMedia, err := h.userSocialMediaService.GetAllUserSocialMedia(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userSocialMedia)
}

func (h *UserSocialMediaHandler) GetUserSocialMediaByID(c *gin.Context) {
	userSocialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user social media ID"})
		return
	}
	userSocialMedia, err := h.userSocialMediaService.GetUserSocialMediaByID(c.Request.Context(), uint(userSocialMediaID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User social media with ID %d not found", userSocialMediaID)})
		return
	}
	c.JSON(http.StatusOK, userSocialMedia)
}

func (h *UserSocialMediaHandler) CreateUserSocialMedia(c *gin.Context) {
	var userSocialMedia model.UserSocialMedia
	if err := c.ShouldBindJSON(&userSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUserSocialMedia, err := h.userSocialMediaService.CreateUserSocialMedia(c.Request.Context(), &userSocialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUserSocialMedia)
}

func (h *UserSocialMediaHandler) UpdateUserSocialMedia(c *gin.Context) {
	userSocialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user social media ID"})
		return
	}
	var updatedUserSocialMedia model.UserSocialMedia
	if err := c.ShouldBindJSON(&updatedUserSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUserSocialMedia.ID = uint(userSocialMediaID)
	userSocialMedia, err := h.userSocialMediaService.UpdateUserSocialMedia(c.Request.Context(), &updatedUserSocialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userSocialMedia)
}

func (h *UserSocialMediaHandler) DeleteUserSocialMedia(c *gin.Context) {
	userSocialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user social media ID"})
		return
	}
	err = h.userSocialMediaService.DeleteUserSocialMedia(c.Request.Context(), uint(userSocialMediaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User social media deleted successfully"})
}
