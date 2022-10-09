package servis

import (
	"context"
	"github.com/google/uuid"
	"tomato/models"
)

type UseCase interface {
	CreateUser(ctx context.Context, user models.CreateUser) (uuid.UUID, error)
	EditUser(ctx context.Context, user models.EditUser) error
	GetUser(ctx context.Context, uuid uuid.UUID) (models.User, error)
	CreateTomato(ctx context.Context, tomato models.CreateTomatoIn) (uuid.UUID, error)
	StartTomato(ctx context.Context, uuid uuid.UUID) error
	GetTomato(ctx context.Context, uuid uuid.UUID) (models.Tomato, error)
	DeleteTomato(ctx context.Context, uuid uuid.UUID) error
	GetTomatoNltx(ctx context.Context, uuid uuid.UUID) (models.TomatoNltx, error)
}
