package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	socialMediaService *service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService *service.SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{socialMediaService}
}

func (h *SocialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	socialMedia, err := h.socialMediaService.GetAllSocialMedia(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, socialMedia)
}

func (h *SocialMediaHandler) GetSocialMediaByID(c *gin.Context) {
	socialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}
	socialMedia, err := h.socialMediaService.GetSocialMediaByID(c.Request.Context(), uint(socialMediaID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Social media with ID %d not found", socialMediaID)})
		return
	}
	c.JSON(http.StatusOK, socialMedia)
}

func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var socialMedia model.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newSocialMedia, err := h.socialMediaService.CreateSocialMedia(c.Request.Context(), &socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newSocialMedia)
}

func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	socialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}
	var updatedSocialMedia model.SocialMedia
	if err := c.ShouldBindJSON(&updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedSocialMedia.ID = int64(socialMediaID)
	socialMedia, err := h.socialMediaService.UpdateSocialMedia(c.Request.Context(), &updatedSocialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, socialMedia)
}

func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	socialMediaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}
	err = h.socialMediaService.DeleteSocialMedia(c.Request.Context(), uint(socialMediaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
