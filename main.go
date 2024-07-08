package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	education2 "github.com/ZoinMe/user-service/handles/education"
	experience2 "github.com/ZoinMe/user-service/handles/experience"
	notification2 "github.com/ZoinMe/user-service/handles/notification"
	skill2 "github.com/ZoinMe/user-service/handles/skill"
	user2 "github.com/ZoinMe/user-service/handles/user"
	userSkill2 "github.com/ZoinMe/user-service/handles/userSkill"
	"github.com/ZoinMe/user-service/services/education"
	"github.com/ZoinMe/user-service/services/experience"
	"github.com/ZoinMe/user-service/services/notification"
	"github.com/ZoinMe/user-service/services/skill"
	"github.com/ZoinMe/user-service/services/user"
	"github.com/ZoinMe/user-service/services/userSkill"
	education3 "github.com/ZoinMe/user-service/stores/education"
	experience3 "github.com/ZoinMe/user-service/stores/experience"
	notification3 "github.com/ZoinMe/user-service/stores/notification"
	skill3 "github.com/ZoinMe/user-service/stores/skill"
	user3 "github.com/ZoinMe/user-service/stores/user"
	userSkill3 "github.com/ZoinMe/user-service/stores/userSkill"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	dbuser := os.Getenv("DB_USER_AIVEN")
	dbpassword := os.Getenv("DB_PASSWORD_AIVEN")
	dbhost := os.Getenv("DB_HOST_AIVEN")
	dbport := os.Getenv("DB_PORT_AIVEN")
	dbdbname := os.Getenv("DB_NAME_AIVEN")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbport, dbdbname)
	// Connect to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow requests from your frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize repositories
	userRepo := user3.NewUserRepository(db)
	skillRepo := skill3.NewSkillRepository(db)
	userSkillRepo := userSkill3.NewUserSkillRepository(db)
	experienceRepo := experience3.NewExperienceRepository(db)
	educationRepo := education3.NewEducationRepository(db)
	notificationRepo := notification3.NewNotificationRepository(db)

	// Initialize services
	userService := user.NewUserService(userRepo)
	skillService := skill.NewSkillService(skillRepo)
	userSkillService := userSkill.NewUserSkillService(userSkillRepo)
	experienceService := experience.NewExperienceService(experienceRepo)
	educationService := education.NewEducationService(educationRepo)
	notificationService := notification.NewNotificationService(notificationRepo)

	// Initialize handlers
	userHandler := user2.NewUserHandler(userService)
	skillHandler := skill2.NewSkillHandler(skillService)
	userSkillHandler := userSkill2.NewUserSkillHandler(userSkillService)
	experienceHandler := experience2.NewExperienceHandler(experienceService)
	educationHandler := education2.NewEducationHandler(educationService)
	notificationHandler := notification2.NewNotificationHandler(notificationService)

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

	router.GET("/user/:id/notification/:id", notificationHandler.GetNotificationByID)
	router.POST("/user/:id/notification", notificationHandler.CreateNotification)
	router.PUT("/user/:id/notification/:id", notificationHandler.UpdateNotification)
	router.DELETE("/user/:id/notification/:id", notificationHandler.DeleteNotification)
	router.GET("/user/:id/notification", notificationHandler.GetNotificationsByUserID)

	// Start the server
	localport := os.Getenv("PORT")

	log.Printf("Server started on port %s", localport)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", localport), router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
