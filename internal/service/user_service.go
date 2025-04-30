package service

import (
	"context"
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Login logs in a user with the given email and password
func (s *UserService) Login(ctx context.Context, email string, password string) (dto.AuthResponse, error) {
	response := dto.AuthResponse{}

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response, err
	}

	claims := &model.JwtCustomClaims{
		int(user.ID),
		user.Name,
		user.Email,
		user.Role.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return response, err
	}

	result := dto.AuthResponse{
		Token: t,
		ID:    int(user.ID),
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role.Name,
	}

	return result, nil
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
