package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
)

type UserSkillHandler struct {
	userSkillService *service.UserSkillService
}

func NewUserSkillHandler(userSkillService *service.UserSkillService) *UserSkillHandler {
	return &UserSkillHandler{userSkillService}
}

func (h *UserSkillHandler) GetAllUserSkills(c *gin.Context) {
	userSkills, err := h.userSkillService.GetAllUserSkills(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userSkills)
}

func (h *UserSkillHandler) GetUserSkillByID(c *gin.Context) {
	userSkillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user skill ID"})
		return
	}
	userSkill, err := h.userSkillService.GetUserSkillByID(c.Request.Context(), uint(userSkillID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User skill with ID %d not found", userSkillID)})
		return
	}
	c.JSON(http.StatusOK, userSkill)
}

func (h *UserSkillHandler) CreateUserSkill(c *gin.Context) {
	var userSkill model.UserSkill
	if err := c.ShouldBindJSON(&userSkill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUserSkill, err := h.userSkillService.CreateUserSkill(c.Request.Context(), &userSkill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUserSkill)
}

func (h *UserSkillHandler) UpdateUserSkill(c *gin.Context) {
	userSkillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user skill ID"})
		return
	}
	var updatedUserSkill model.UserSkill
	if err := c.ShouldBindJSON(&updatedUserSkill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUserSkill.ID = int64(userSkillID)
	userSkill, err := h.userSkillService.UpdateUserSkill(c.Request.Context(), &updatedUserSkill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userSkill)
}

func (h *UserSkillHandler) DeleteUserSkill(c *gin.Context) {
	userSkillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user skill ID"})
		return
	}
	err = h.userSkillService.DeleteUserSkill(c.Request.Context(), uint(userSkillID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User skill deleted successfully"})
}

func (ush *UserSkillHandler) GetUserSkillsByUserID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userSkills, err := ush.userSkillService.GetUserSkillsByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userSkills)
}
