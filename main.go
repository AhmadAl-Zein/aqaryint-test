package main

import (
	"log"
	"os"
	"strconv"

	"github.com/AhmadAl-Zein/aqary_test/controller"
	"github.com/AhmadAl-Zein/aqary_test/database"
	"github.com/AhmadAl-Zein/aqary_test/db"
	"github.com/AhmadAl-Zein/aqary_test/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	userService    service.UserService
	userController controller.UserController
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	postgres, err := db.NewPostgres(os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"), dbPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	queries := database.New(postgres.DB)
	userService = service.New(queries)
	userController = controller.New(userService)

	r := gin.Default()

	userRoutes := r.Group("api/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/generateotp", userController.GenerateOTP)
		userRoutes.POST("/verifyotp", userController.VerifyOTP)
	}

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
