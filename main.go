package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ZoinMe/user-service/handler"
	"github.com/ZoinMe/user-service/repository"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "root:ganesh123@tcp/zoinme?parseTime=true")

	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	skillRepo := repository.NewSkillRepository(db)
	userSkillRepo := repository.NewUserSkillRepository(db)
	experienceRepo := repository.NewExperienceRepository(db)
	educationRepo := repository.NewEducationRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	skillService := service.NewSkillService(skillRepo)
	userSkillService := service.NewUserSkillService(userSkillRepo)
	experienceService := service.NewExperienceService(experienceRepo)
	educationService := service.NewEducationService(educationRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	skillHandler := handler.NewSkillHandler(skillService)
	userSkillHandler := handler.NewUserSkillHandler(userSkillService)
	experienceHandler := handler.NewExperienceHandler(experienceService)
	educationHandler := handler.NewEducationHandler(educationService)

	// Define APIs for each entity
	router.GET("/user", userHandler.GetUsers)
	router.GET("/user/:id", userHandler.GetUserByID)
	router.POST("/user", userHandler.CreateUser)
	router.PUT("/user/:id", userHandler.UpdateUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)

	router.GET("/skill", skillHandler.GetAllSkills)
	router.GET("/skill/:id", skillHandler.GetSkillByID)
	router.POST("/skill", skillHandler.CreateSkill)
	router.PUT("/skill/:id", skillHandler.UpdateSkill)
	router.DELETE("/skill/:id", skillHandler.DeleteSkill)

	router.GET("/user/:id/skill/:id", userSkillHandler.GetUserSkillByID)
	router.POST("/user/:id/skill", userSkillHandler.CreateUserSkill)
	router.PUT("/user/:id/skill/:id", userSkillHandler.UpdateUserSkill)
	router.DELETE("/user/:id/skill/:id", userSkillHandler.DeleteUserSkill)
	router.GET("/user/:id/skill", userSkillHandler.GetUserSkillsByUserID)

	router.GET("/experience", experienceHandler.GetAllExperiences)
	router.GET("/experience/:id", experienceHandler.GetExperienceByID)
	router.POST("/experience", experienceHandler.CreateExperience)
	router.PUT("/experience/:id", experienceHandler.UpdateExperience)
	router.DELETE("/experience/:id", experienceHandler.DeleteExperience)
	router.GET("/user/:id/experience", experienceHandler.GetExperiencesByUserID)

	router.GET("/education", educationHandler.GetAllEducations)
	router.GET("/education/:id", educationHandler.GetEducationByID)
	router.POST("/education", educationHandler.CreateEducation)
	router.PUT("/education/:id", educationHandler.UpdateEducation)
	router.DELETE("/education/:id", educationHandler.DeleteEducation)
	router.GET("/user/:id/education", educationHandler.GetEducationsByUserID)

	// Start the server
	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
