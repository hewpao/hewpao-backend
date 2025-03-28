package repository

import (
	"context"

	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/dto"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, req dto.CreateUserDTO) error
	UpdateVerification(ctx context.Context, req *domain.User) error
	EditProfile(ctx context.Context, userID string, req dto.EditProfileDTO) error
}
