package main

import (
	"fmt"
	"jagratama-backend/internal/handler"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"jagratama-backend/internal/service"

	customMiddleware "jagratama-backend/internal/middleware"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
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
	err = db.AutoMigrate(&model.User{}, &model.Document{}, &model.ApprovalRequest{}, &model.RefreshToken{})
	if err != nil {
		fmt.Printf("Failed to migrate database %v", err)
	}

	fmt.Println("Successfully connected to database")

	refreshTokenRepository := repository.NewRefreshTokenRepository(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository, *refreshTokenRepository)
	userHandler := handler.NewUserHandler(*userService)

	approvalRequestRepository := repository.NewApprovalRequestRepository(db)
	documentRepository := repository.NewDocumentRepository(db)
	documentService := service.NewDocumentService(*documentRepository, *approvalRequestRepository, *userRepository)
	documentHandler := handler.NewDocumentHandler(*documentService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Route
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/login", userHandler.Login)

		v1WithAuth := v1.Group("")
		v1WithAuth.Use(echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte("secret"),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(model.JwtCustomClaims)
			},
		}))
		v1WithAuth.Use(customMiddleware.Auth)

		v1WithAuth.GET("/auth/me", userHandler.GetMe)
		v1WithAuth.POST("/auth/refresh-token", userHandler.RefreshToken)
		v1WithAuth.POST("/auth/logout", userHandler.Logout)

		v1WithAuth.GET("/users", userHandler.GetAllUsers)
		v1WithAuth.POST("/users", userHandler.CreateUser)
		v1WithAuth.PUT("/users/profile", userHandler.UpdateUserProfile)
		v1WithAuth.GET("/users/approver-reviewer", userHandler.GetApproverReviewerUsers)
		v1WithAuth.GET("/users/:id", userHandler.GetUserByID)
		v1WithAuth.PUT("/users/:id", userHandler.UpdateUser)
		v1WithAuth.DELETE("/users/:id", userHandler.DeleteUser)

		v1WithAuth.GET("/documents", documentHandler.GetAllDocuments)
		v1WithAuth.GET("/documents/:slug", documentHandler.GetDocumentBySlug)
		v1WithAuth.POST("/documents", documentHandler.CreateDocument, customMiddleware.RoleCheck([]string{"requester", "admin"}))
		v1WithAuth.PUT("/documents/:slug", documentHandler.UpdateDocument)
		v1WithAuth.DELETE("/documents/:slug", documentHandler.DeleteDocument)

		v1WithAuth.GET("/documents/:slug/tracking", documentHandler.GetDocumentProgress)
		v1WithAuth.POST("/documents/:slug/approval", documentHandler.ApprovalAction)

		v1WithAuth.GET("/documents/to-review", documentHandler.GetDocumentApprovalRequest)
		v1WithAuth.GET("/documents/to-review/history", documentHandler.GetDocumentApprovalHistory)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
