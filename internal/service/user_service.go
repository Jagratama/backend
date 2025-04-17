package service

import (
	"context"
	"fmt"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Login logs in a user with the given email and password
func (s *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	fmt.Println(user.Password)
	fmt.Println(password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	claims := &JwtCustomClaims{
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
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
	fmt.Println(user.Password)
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
