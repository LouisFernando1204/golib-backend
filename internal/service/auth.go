package service

import (
	"context"
	"errors"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(conf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		conf:           conf,
		userRepository: userRepository,
	}
}

// Login implements domain.AuthService.
func (a *authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("user password doesn't match")
	}

	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("jwt signing failed")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
