package service

import (
	"errors"

	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"github.com/rheadavin/hr-go-api/internal/repository"
	"github.com/rheadavin/hr-go-api/pkg/hash"
	"github.com/rheadavin/hr-go-api/pkg/jwt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*models.User, error) {
	// cek email sudah ada
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// hash password
	hashedPw, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPw,
		Role:     "staff",
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	// get user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email or password not valid")
	}

	// compare password
	if !hash.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("email or password not valid")
	}

	// generate token jwt
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserProfile{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
