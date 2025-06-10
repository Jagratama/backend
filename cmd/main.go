package main

import (
	"fmt"
	"jagratama-backend/internal/config"
	"jagratama-backend/internal/database/seeder"
	"jagratama-backend/internal/handler"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/pkg/aws"
	"jagratama-backend/internal/repository"
	"jagratama-backend/internal/service"

	customMiddleware "jagratama-backend/internal/middleware"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.SetupEnv()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Printf("Failed to connect to database %v", err)
		return
	}
	fmt.Println("Successfully connected to database")

	// Auto migrate the User table
	err = config.MigrateDB(db)
	if err != nil {
		fmt.Printf("Failed to migrate database %v", err)
		return
	}
	fmt.Println("Successfully migrated database")

	if err := seeder.RunAll(db); err != nil {
		fmt.Printf("Failed to run seeders: %v", err)
	} else {
		fmt.Println("Successfully seeded database")
	}

	s3Uploader, err := aws.NewS3Uploader(config.GetEnv("AWS_BUCKET_NAME", ""))
	if err != nil {
		fmt.Printf("Failed to create S3 uploader: %v", err)
		return
	}

	fileRepository := repository.NewFileRepository(db)
	refreshTokenRepository := repository.NewRefreshTokenRepository(db)
	userRepository := repository.NewUserRepository(db)
	approvalRequestRepository := repository.NewApprovalRequestRepository(db)
	documentRepository := repository.NewDocumentRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	positionCategoryRuleRepository := repository.NewPositionCategoryRuleRepository(db)
	positionRepository := repository.NewPositionRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	fileService := service.NewFileService(*fileRepository, s3Uploader)
	userService := service.NewUserService(*userRepository, *refreshTokenRepository)
	documentService := service.NewDocumentService(*documentRepository, *approvalRequestRepository, *userRepository, *positionCategoryRuleRepository)
	roleService := service.NewRoleService(*roleRepository)
	positionService := service.NewPositionService(*positionRepository, *positionCategoryRuleRepository)
	categoryService := service.NewCategoryService(*categoryRepository)

	fileHandler := handler.NewFileHandler(*fileService)
	userHandler := handler.NewUserHandler(*userService)
	documentHandler := handler.NewDocumentHandler(*documentService)
	roleHandler := handler.NewRoleHandler(*roleService)
	positionHandler := handler.NewPositionHandler(*positionService)
	categoryHandler := handler.NewCategoryHandler(*categoryService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Route
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/login", userHandler.Login)
		v1.POST("/auth/refresh-token", userHandler.RefreshToken)

		v1WithAuth := v1.Group("")
		v1WithAuth.Use(echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte(config.GetEnv("JWT_ACCESS_TOKEN_SECRET", "secret")),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(model.JwtCustomClaims)
			},
		}))
		v1WithAuth.Use(customMiddleware.Auth)

		v1WithAuth.GET("/auth/me", userHandler.GetMe)
		v1WithAuth.POST("/auth/logout", userHandler.Logout)

		v1WithAuth.POST("/upload", fileHandler.UploadFile)

		v1WithAuth.GET("/users", userHandler.GetAllUsers)
		v1WithAuth.POST("/users", userHandler.CreateUser)
		v1WithAuth.PUT("/users/profile", userHandler.UpdateUserProfile)
		v1WithAuth.GET("/users/approver-reviewer", userHandler.GetApproverReviewerUsers)
		v1WithAuth.GET("/users/:id", userHandler.GetUserByID)
		v1WithAuth.PUT("/users/:id", userHandler.UpdateUser)
		v1WithAuth.DELETE("/users/:id", userHandler.DeleteUser)

		v1WithAuth.GET("/roles", roleHandler.GetAllRoles)
		v1WithAuth.GET("/roles/:id", roleHandler.GetRoleByID)

		v1WithAuth.GET("/positions", positionHandler.GetAllPositions)
		v1WithAuth.GET("/positions/:id", positionHandler.GetPositionByID)
		v1WithAuth.GET("/positions/rules-by-category/:categoryID", positionHandler.GetPositionsRequiredByCategoryID)

		v1WithAuth.GET("/documents", documentHandler.GetAllDocuments)
		v1WithAuth.GET("/documents/:slug", documentHandler.GetDocumentBySlug)
		v1WithAuth.POST("/documents", documentHandler.CreateDocument, customMiddleware.RoleCheck([]string{"requester"}))
		v1WithAuth.PUT("/documents/:slug", documentHandler.UpdateDocument)
		v1WithAuth.DELETE("/documents/:slug", documentHandler.DeleteDocument)

		v1WithAuth.POST("/documents/:slug/confirm", documentHandler.ConfirmDocument)
		v1WithAuth.POST("/documents/:slug/reupload/:approvalID", documentHandler.ReuploadDocument)
		v1WithAuth.GET("/documents/:slug/tracking", documentHandler.GetDocumentProgress)
		v1WithAuth.POST("/documents/:slug/approval", documentHandler.ApprovalAction)

		v1WithAuth.GET("/documents/to-review", documentHandler.GetDocumentApprovalRequest)
		v1WithAuth.GET("/documents/to-review/history", documentHandler.GetDocumentApprovalHistory)
		v1WithAuth.GET("/documents/to-review/:slug", documentHandler.GetDocumentApprovalReviewDetail)

		v1WithAuth.GET("/documents/counter", documentHandler.GetCountAllMyDocuments)

		v1WithAuth.GET("/categories", categoryHandler.GetAllCategories)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
