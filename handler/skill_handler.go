package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	skillService *service.SkillService
}

func NewSkillHandler(skillService *service.SkillService) *SkillHandler {
	return &SkillHandler{skillService}
}

func (h *SkillHandler) GetAllSkills(c *gin.Context) {
	skills, err := h.skillService.GetAllSkills(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}

func (h *SkillHandler) GetSkillByID(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	skill, err := h.skillService.GetSkillByID(c.Request.Context(), uint(skillID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Skill with ID %d not found", skillID)})
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	var skill model.Skill

	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSkill, err := h.skillService.CreateSkill(c.Request.Context(), &skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSkill)
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	var updatedSkill model.Skill

	if err := c.ShouldBindJSON(&updatedSkill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSkill.ID = int64(skillID)

	skill, err := h.skillService.UpdateSkill(c.Request.Context(), &updatedSkill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	err = h.skillService.DeleteSkill(c.Request.Context(), uint(skillID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}
