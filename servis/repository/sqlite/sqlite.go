package sqlite

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"tomato/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {

	return uuid.Nil, nil
}
func (r *Repository) EditUser(ctx context.Context, user models.User) error {
	return nil
}
func (r *Repository) GetUser(ctx context.Context, uuid uuid.UUID) (models.User, error) {
	return models.User{}, nil
}
func (r *Repository) CreateTomato(ctx context.Context, tomato models.Tomato) (uuid.UUID, error) {
	return uuid.Nil, nil
}
func (r *Repository) CompleteTomato(ctx context.Context, uuid uuid.UUID) error {
	return nil
}
func (r *Repository) GetTomato(ctx context.Context, uuid uuid.UUID) (models.Tomato, error) {
	return models.Tomato{}, nil
}
func (r *Repository) DeleteTomato(ctx context.Context, uuid uuid.UUID) error {
	return nil
}
func (r *Repository) GetTomatoNltx(ctx context.Context, uuid2 uuid.UUID) (models.TomatoNltx, error) {
	return models.TomatoNltx{}, nil
}
