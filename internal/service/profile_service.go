package service

import (
	"context"
	"time"

	"sea-cinema-api/internal/contract"
	"sea-cinema-api/internal/model"
	"sea-cinema-api/internal/repository"
)

type ProfileService interface {
	GetProfile(ctx context.Context, claims model.JWTClaims) (contract.GetProfileResponse, error)
	TopUpBalance(ctx context.Context, claims model.JWTClaims, request contract.TopUpBalanceRequest) (contract.TopUpBalanceResponse, error)
}

type profileService struct {
	userRepository repository.UserRepository
}

func NewProfileService(userRepository repository.UserRepository) ProfileService {
	return &profileService{
		userRepository: userRepository,
	}
}

func (service *profileService) GetProfile(ctx context.Context, claims model.JWTClaims) (contract.GetProfileResponse, error) {
	user, err := service.userRepository.GetUserById(ctx, claims.User.Id)
	if err != nil {
		return contract.GetProfileResponse{}, err
	}

	response := contract.GetProfileResponse{
		Id:        user.Id,
		Username:  user.Username,
		Name:      user.Name,
		Balance:   user.Balance,
		BirthDate: user.BirthDate,
	}

	return response, nil
}

func (service *profileService) TopUpBalance(ctx context.Context, claims model.JWTClaims, request contract.TopUpBalanceRequest) (contract.TopUpBalanceResponse, error) {
	user, err := service.userRepository.GetUserById(ctx, claims.User.Id)
	if err != nil {
		return contract.TopUpBalanceResponse{}, err
	}

	user.Balance += request.Balance
	user.UpdatedAt = time.Now()

	if err := service.userRepository.UpdateUser(ctx, claims.User.Id, user); err != nil {
		return contract.TopUpBalanceResponse{}, err
	}

	response := contract.TopUpBalanceResponse{
		Id:       user.Id,
		Username: user.Username,
		Balance:  user.Balance,
	}

	return response, nil
}
