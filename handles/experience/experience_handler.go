package experience

import (
	"fmt"
	"github.com/ZoinMe/user-service/services"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	experienceService services.Experience
}

func NewExperienceHandler(experienceService services.Experience) *ExperienceHandler {
	return &ExperienceHandler{experienceService}
}

func (h *ExperienceHandler) GetAllExperiences(c *gin.Context) {
	experiences, err := h.experienceService.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

func (h *ExperienceHandler) GetExperienceByID(c *gin.Context) {
	experienceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experience ID"})
		return
	}

	experience, err := h.experienceService.GetByID(c.Request.Context(), uint(experienceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Experience with ID %d not found", experienceID)})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func (h *ExperienceHandler) CreateExperience(c *gin.Context) {
	var experience model.Experience

	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newExperience, err := h.experienceService.Create(c.Request.Context(), &experience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newExperience)
}

func (h *ExperienceHandler) UpdateExperience(c *gin.Context) {
	experienceID := c.Param("id")

	var updatedExperience model.Experience
	if err := c.ShouldBindJSON(&updatedExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedExperience.ID = experienceID

	experience, err := h.experienceService.Update(c.Request.Context(), &updatedExperience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func (h *ExperienceHandler) DeleteExperience(c *gin.Context) {
	experienceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experience ID"})
		return
	}

	err = h.experienceService.Delete(c.Request.Context(), uint(experienceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}

func (eh *ExperienceHandler) GetExperiencesByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	experiences, err := eh.experienceService.GetByUserID(c.Request.Context(), userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)
}
