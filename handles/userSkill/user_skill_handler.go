package userSkill

import (
	"fmt"
	"github.com/ZoinMe/user-service/services"
	"net/http"
	"strconv"

	"github.com/ZoinMe/user-service/model"
	"github.com/gin-gonic/gin"
)

type UserSkillHandler struct {
	userSkillService services.UserSkill
}

func NewUserSkillHandler(userSkillService services.UserSkill) *UserSkillHandler {
	return &UserSkillHandler{userSkillService}
}

func (h *UserSkillHandler) GetAllUserSkills(c *gin.Context) {
	userSkills, err := h.userSkillService.Get(c.Request.Context())
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

	userSkill, err := h.userSkillService.GetByID(c.Request.Context(), uint(userSkillID))
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

	newUserSkill, err := h.userSkillService.Create(c.Request.Context(), &userSkill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUserSkill)
}

func (h *UserSkillHandler) UpdateUserSkill(c *gin.Context) {
	userSkillID:=c.Param("id")
	var updatedUserSkill model.UserSkill

	if err := c.ShouldBindJSON(&updatedUserSkill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUserSkill.ID = userSkillID

	userSkill, err := h.userSkillService.Update(c.Request.Context(), &updatedUserSkill)
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

	err = h.userSkillService.Delete(c.Request.Context(), uint(userSkillID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User skill deleted successfully"})
}

func (ush *UserSkillHandler) GetUserSkillsByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	userSkills, err := ush.userSkillService.GetByUserID(c.Request.Context(), userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userSkills)
}
