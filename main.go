package main

import (
	"log"
	"net/http"

	"github.com/ZoinMe/user-service/handler"
	"github.com/ZoinMe/user-service/model"
	"github.com/ZoinMe/user-service/repository"
	"github.com/ZoinMe/user-service/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Connect to the database
	db, err := gorm.Open("mysql", "root:ganesh123@/zoinme?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Auto migrate the database
	db.AutoMigrate(&model.User{})

	db.SingularTable(true)

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
router.GET("/users", userHandler.GetUsers)
router.GET("/users/:id", userHandler.GetUserByID)
router.POST("/users", userHandler.CreateUser)
router.PUT("/users/:id", userHandler.UpdateUser)
router.DELETE("/users/:id", userHandler.DeleteUser)

router.GET("/skills", skillHandler.GetAllSkills)
router.GET("/skills/:id", skillHandler.GetSkillByID)
router.POST("/skills", skillHandler.CreateSkill)
router.PUT("/skills/:id", skillHandler.UpdateSkill)
router.DELETE("/skills/:id", skillHandler.DeleteSkill)

router.GET("/users/:userID/skills", userSkillHandler.GetAllUserSkills)
router.GET("/users/:userID/skills/:id", userSkillHandler.GetUserSkillByID)
router.POST("/users/:userID/skills", userSkillHandler.CreateUserSkill)
router.PUT("/users/:userID/skills/:id", userSkillHandler.UpdateUserSkill)
router.DELETE("/users/:userID/skills/:skillID", userSkillHandler.DeleteUserSkill)

router.GET("/socialmedia", socialMediaHandler.GetAllSocialMedia)
router.GET("/socialmedia/:id", socialMediaHandler.GetSocialMediaByID)
router.POST("/socialmedia", socialMediaHandler.CreateSocialMedia)
router.PUT("/socialmedia/:id", socialMediaHandler.UpdateSocialMedia)
router.DELETE("/socialmedia/:id", socialMediaHandler.DeleteSocialMedia)

router.GET("/users/:userID/socialmedia", userSocialMediaHandler.GetAllUserSocialMedia)
router.GET("/users/:userID/socialmedia/:id", userSocialMediaHandler.GetUserSocialMediaByID)
router.POST("/users/:userID/socialmedia", userSocialMediaHandler.CreateUserSocialMedia)
router.PUT("/users/:userID/socialmedia/:id", userSocialMediaHandler.UpdateUserSocialMedia)
router.DELETE("/users/:userID/socialmedia/:socialMediaID", userSocialMediaHandler.DeleteUserSocialMedia)

router.GET("/experiences", experienceHandler.GetAllExperiences)
router.GET("/experiences/:id", experienceHandler.GetExperienceByID)
router.POST("/experiences", experienceHandler.CreateExperience)
router.PUT("/experiences/:id", experienceHandler.UpdateExperience)
router.DELETE("/experiences/:id", experienceHandler.DeleteExperience)

router.GET("/educations", educationHandler.GetAllEducations)
router.GET("/educations/:id", educationHandler.GetEducationByID)
router.POST("/educations", educationHandler.CreateEducation)
router.PUT("/educations/:id", educationHandler.UpdateEducation)
router.DELETE("/educations/:id", educationHandler.DeleteEducation)


	// Start the server
	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
