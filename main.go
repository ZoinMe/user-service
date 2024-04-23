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

	db, err := sql.Open("mysql", "root:ganesh123@/zoinme")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	skillRepo := repository.NewSkillRepository(db)
	userSkillRepo := repository.NewUserSkillRepository(db)
	socialMediaRepo := repository.NewSocialMediaRepository(db)
	userSocialMediaRepo := repository.NewUserSocialMediaRepository(db)
	experienceRepo := repository.NewExperienceRepository(db)
	educationRepo := repository.NewEducationRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	skillService := service.NewSkillService(skillRepo)
	userSkillService := service.NewUserSkillService(userSkillRepo)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	userSocialMediaService := service.NewUserSocialMediaService(userSocialMediaRepo)
	experienceService := service.NewExperienceService(experienceRepo)
	educationService := service.NewEducationService(educationRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	skillHandler := handler.NewSkillHandler(skillService)
	userSkillHandler := handler.NewUserSkillHandler(userSkillService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)
	userSocialMediaHandler := handler.NewUserSocialMediaHandler(userSocialMediaService)
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

	router.GET("/user/:id/skill", userSkillHandler.GetAllUserSkills)
	router.GET("/user/:id/skill/:id", userSkillHandler.GetUserSkillByID)
	router.POST("/user/:id/skill", userSkillHandler.CreateUserSkill)
	router.PUT("/user/:id/skill/:id", userSkillHandler.UpdateUserSkill)
	router.DELETE("/user/:id/skill/:id", userSkillHandler.DeleteUserSkill)

	router.GET("/socialmedia", socialMediaHandler.GetAllSocialMedia)
	router.GET("/socialmedia/:id", socialMediaHandler.GetSocialMediaByID)
	router.POST("/socialmedia", socialMediaHandler.CreateSocialMedia)
	router.PUT("/socialmedia/:id", socialMediaHandler.UpdateSocialMedia)
	router.DELETE("/socialmedia/:id", socialMediaHandler.DeleteSocialMedia)

	router.GET("/user/:id/socialmedia", userSocialMediaHandler.GetAllUserSocialMedia)
	router.GET("/user/:id/socialmedia/:id", userSocialMediaHandler.GetUserSocialMediaByID)
	router.POST("/user/:id/socialmedia", userSocialMediaHandler.CreateUserSocialMedia)
	router.PUT("/user/:id/socialmedia/:id", userSocialMediaHandler.UpdateUserSocialMedia)
	router.DELETE("/user/:id/socialmedia/:socialMediaID", userSocialMediaHandler.DeleteUserSocialMedia)

	router.GET("/experience", experienceHandler.GetAllExperiences)
	router.GET("/experience/:id", experienceHandler.GetExperienceByID)
	router.POST("/experience", experienceHandler.CreateExperience)
	router.PUT("/experience/:id", experienceHandler.UpdateExperience)
	router.DELETE("/experience/:id", experienceHandler.DeleteExperience)

	router.GET("/education", educationHandler.GetAllEducations)
	router.GET("/education/:id", educationHandler.GetEducationByID)
	router.POST("/education", educationHandler.CreateEducation)
	router.PUT("/education/:id", educationHandler.UpdateEducation)
	router.DELETE("/education/:id", educationHandler.DeleteEducation)

	// Start the server
	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
