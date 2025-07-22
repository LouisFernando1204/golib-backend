package domain

import (
	"context"

	"github.com/LouisFernando1204/golang-restapi.git/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
