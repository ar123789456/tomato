package servis

import (
	"context"
	"github.com/google/uuid"
	"tomato/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
	EditUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, uuid uuid.UUID) (models.User, error)
	CreateTomato(ctx context.Context, tomato models.Tomato) (uuid.UUID, error)
	CompleteTomato(ctx context.Context, uuid uuid.UUID) error
	GetTomato(ctx context.Context, uuid uuid.UUID) (models.Tomato, error)
	DeleteTomato(ctx context.Context, uuid uuid.UUID) error
	GetTomatoNltx(ctx context.Context, uuid2 uuid.UUID) (models.TomatoNltx, error)
}
