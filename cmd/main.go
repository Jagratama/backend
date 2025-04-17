package main

import (
	"fmt"
	"jagratama-backend/internal/handler"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"jagratama-backend/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	helpers.SetupConfig()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", helpers.GetEnv("DB_HOST", "localhost"), helpers.GetEnv("DB_USER", ""), helpers.GetEnv("DB_NAME", ""), helpers.GetEnv("DB_PORT", ""))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate the User table
	if err := db.AutoMigrate(&model.User{}); err != nil {
		fmt.Printf("Failed to migrate database %v", err)
	}

	fmt.Println("Successfully connected to database")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(echojwt.JWT([]byte("secret")))

	// Route
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/login", userHandler.Login)

		v1.GET("/users", userHandler.GetAllUsers)
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUserByID)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
