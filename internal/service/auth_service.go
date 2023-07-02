package service

import (
	"context"

	"sea-cinema-api/internal/contract"
	"sea-cinema-api/internal/model"
	"sea-cinema-api/internal/repository"
	"sea-cinema-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, request contract.RegisterRequest) (contract.RegisterResponse, error)
	Login(ctx context.Context, request contract.LoginRequest) (contract.LoginResponse, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) Register(ctx context.Context, request contract.RegisterRequest) (contract.RegisterResponse, error) {
	// check if the username already exists
	_, err := service.userRepository.GetUserByUsername(ctx, request.Username)
	if err == nil {
		return contract.RegisterResponse{}, fiber.NewError(fiber.StatusConflict, "username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return contract.RegisterResponse{}, err
	}

	user := model.User{
		Username:  request.Username,
		Password:  string(hashedPassword),
		Name:      request.Name,
		BirthDate: request.BirthDate,
	}

	err = service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return contract.RegisterResponse{}, err
	}

	// generate access token
	userClaimsData := model.JWTUserClaimsData{
		Username:  request.Username,
		Name:      request.Name,
		BirthDate: request.BirthDate,
	}

	accessToken, err := utils.GenerateJwt(userClaimsData)
	if err != nil {
		return contract.RegisterResponse{}, err
	}

	response := contract.RegisterResponse{
		AccessToken: accessToken,
	}
	return response, nil
}

func (service *authService) Login(ctx context.Context, request contract.LoginRequest) (contract.LoginResponse, error) {
	user, err := service.userRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return contract.LoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return contract.LoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	// generate access token
	userClaimsData := model.JWTUserClaimsData{
		Username:  user.Username,
		Name:      user.Name,
		BirthDate: user.BirthDate,
	}

	accessToken, err := utils.GenerateJwt(userClaimsData)
	if err != nil {
		return contract.LoginResponse{}, err
	}

	response := contract.LoginResponse{
		AccessToken: accessToken,
	}
	return response, nil
}