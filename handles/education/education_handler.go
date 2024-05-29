package education

import (
	"fmt"
	"github.com/ZoinMe/user-service/services"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/gin-gonic/gin"
)

type EducationHandler struct {
	educationService services.Education
}

func NewEducationHandler(educationService services.Education) *EducationHandler {
	return &EducationHandler{educationService}
}

func (h *EducationHandler) GetAllEducations(c *gin.Context) {
	educations, err := h.educationService.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, educations)
}

func (h *EducationHandler) GetEducationByID(c *gin.Context) {
	educationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid education ID"})
		return
	}

	education, err := h.educationService.GetByID(c.Request.Context(), uint(educationID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Education with ID %d not found", educationID)})
		return
	}

	c.JSON(http.StatusOK, education)
}

func (h *EducationHandler) CreateEducation(c *gin.Context) {
	var education model.Education

	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEducation, err := h.educationService.Create(c.Request.Context(), &education)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEducation)
}

func (h *EducationHandler) UpdateEducation(c *gin.Context) {
	educationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid education ID"})
		return
	}

	var updatedEducation model.Education

	if err := c.ShouldBindJSON(&updatedEducation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEducation.ID = int64(educationID)

	education, err := h.educationService.Update(c.Request.Context(), &updatedEducation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, education)
}

func (h *EducationHandler) DeleteEducation(c *gin.Context) {
	educationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid education ID"})
		return
	}

	err = h.educationService.Delete(c.Request.Context(), uint(educationID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education deleted successfully"})
}

func (eh *EducationHandler) GetEducationsByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	educations, err := eh.educationService.GetByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, educations)
}
