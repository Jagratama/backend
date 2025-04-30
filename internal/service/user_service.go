package service

import (
	"context"
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewUserService(userRepository repository.UserRepository, refreshTokenRepository repository.RefreshTokenRepository) *UserService {
	return &UserService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

// Login logs in a user with the given email and password
func (s *UserService) Login(ctx context.Context, email string, password string) (*dto.AuthResponse, error) {
	response := &dto.AuthResponse{}

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response, err
	}

	accessTokenTime := helpers.GetEnv("JWT_ACCESS_TOKEN_EXPIRES", "3600")
	accessTokenTimeInt, err := strconv.Atoi(accessTokenTime)
	if err != nil {
		return response, err
	}

	jwtExpireTime := jwt.NewNumericDate(time.Now().Add(time.Duration(accessTokenTimeInt) * time.Second))

	claims := &model.JwtCustomClaims{
		int(user.ID),
		user.Name,
		user.Email,
		user.Role.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwtExpireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(helpers.GetEnv("JWT_ACCESS_TOKEN_SECRET", "secret")))
	if err != nil {
		return response, err
	}

	RefreshTokenTime := helpers.GetEnv("JWT_REFRESH_TOKEN_EXPIRES", "604800")
	RefreshTokenTimeInt, err := strconv.Atoi(RefreshTokenTime)
	if err != nil {
		return response, err
	}
	RefreshTokenExpireTime := jwt.NewNumericDate(time.Now().Add(time.Duration(RefreshTokenTimeInt) * time.Second))
	claimsRefresh := &model.JwtCustomClaims{
		int(user.ID),
		user.Name,
		user.Email,
		user.Role.Name,
		jwt.RegisteredClaims{
			ExpiresAt: RefreshTokenExpireTime,
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	RefreshToken, err := refreshToken.SignedString([]byte(helpers.GetEnv("JWT_REFRESH_TOKEN_SECRET", "secret")))
	if err != nil {
		return response, err
	}

	refreshTokenData := &model.RefreshToken{
		UserID:    int(user.ID),
		Token:     RefreshToken,
		UserAgent: "",
		ExpiredAt: RefreshTokenExpireTime.String(),
	}

	err = s.refreshTokenRepository.Create(ctx, refreshTokenData)
	if err != nil {
		return response, err
	}

	result := dto.AuthResponse{
		ID:           int(user.ID),
		Email:        user.Email,
		Name:         user.Name,
		Role:         user.Role.Name,
		Token:        t,
		RefreshToken: RefreshToken,
	}

	return &result, nil
}

// GetAllUsers retrieves all users from the database
func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		user.Password = ""
	}

	return users, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	newUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	newUser.Password = ""
	return newUser, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// Check if the user exists
	existingUser, err := s.userRepository.GetUserByID(ctx, int(user.ID))
	if err != nil {
		return nil, err
	}

	user.Password = existingUser.Password

	// Save the updated user to the database
	updatedUser, err := s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	updatedUser.Password = ""
	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	_, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	user, err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	user.Password = ""
	return nil
}

func (s *UserService) GetMe(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		ImagePath:  user.ImagePath,
		RoleID:     user.RoleID,
		PositionID: user.PositionID,
		Role: dto.Role{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
		Position: dto.Position{
			ID:                 user.Position.ID,
			Name:               user.Position.Name,
			RequiresSignatures: user.Position.RequiresSignatures,
		},
	}
	return response, nil
}

func (s *UserService) GetApproverReviewerUsers(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := s.userRepository.GetApproverReviewerUsers(ctx)
	if err != nil {
		return nil, err
	}

	response := []*dto.UserResponse{}

	for _, user := range users {
		response = append(response, &dto.UserResponse{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			ImagePath:  user.ImagePath,
			RoleID:     user.RoleID,
			PositionID: user.PositionID,
			Role: dto.Role{
				ID:   user.Role.ID,
				Name: user.Role.Name,
			},
			Position: dto.Position{
				ID:                 user.Position.ID,
				Name:               user.Position.Name,
				RequiresSignatures: user.Position.RequiresSignatures,
			},
		})
	}

	return response, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, user *dto.UpdateProfileRequest, userID int) (*dto.UserResponse, error) {
	// Check if the user exists
	existingUser, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Update the user fields
	existingUser.Name = user.Name
	existingUser.ImagePath = user.ImagePath

	// Save the updated user to the database
	updatedUser, err := s.userRepository.UpdateUser(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:         updatedUser.ID,
		Name:       updatedUser.Name,
		Email:      updatedUser.Email,
		ImagePath:  updatedUser.ImagePath,
		RoleID:     updatedUser.RoleID,
		PositionID: updatedUser.PositionID,
	}

	return response, nil
}

func (s *UserService) RefreshToken(ctx context.Context, userID int, refreshToken string) (*dto.AuthResponse, error) {
	response := &dto.AuthResponse{}

	// Check if the user exists
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return response, err
	}
	// Check if the refresh token is valid
	refreshTokenData, err := s.refreshTokenRepository.GetByUserID(ctx, userID)
	if err != nil {
		return response, err
	}
	if refreshTokenData.Token != refreshToken {
		return response, err
	}
	// Parse the refresh token
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(helpers.GetEnv("JWT_REFRESH_TOKEN_SECRET", "secret")), nil
	})
	if err != nil {
		return response, err
	}
	// Check if the token is valid
	if _, ok := parsedToken.Claims.(*model.JwtCustomClaims); ok && parsedToken.Valid {
		// Create a new access token
		accessTokenTime := helpers.GetEnv("JWT_ACCESS_TOKEN_EXPIRES", "3600")
		accessTokenTimeInt, err := strconv.Atoi(accessTokenTime)
		if err != nil {
			return response, err
		}
		jwtExpireTime := jwt.NewNumericDate(time.Now().Add(time.Duration(accessTokenTimeInt) * time.Second))

		newClaims := &model.JwtCustomClaims{
			int(user.ID),
			user.Name,
			user.Email,
			user.Role.Name,
			jwt.RegisteredClaims{
				ExpiresAt: jwtExpireTime,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		t, err := token.SignedString([]byte(helpers.GetEnv("JWT_ACCESS_TOKEN_SECRET", "secret")))
		if err != nil {
			return response, err
		}

		response.ID = int(user.ID)
		response.Email = user.Email
		response.Name = user.Name
		response.Role = user.Role.Name
		response.Token = t
		response.RefreshToken = refreshTokenData.Token

		return response, nil
	}
	// If the token is not valid, return an error
	return response, err
}
