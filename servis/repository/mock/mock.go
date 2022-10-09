package mock

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"tomato/models"
)

type Repository struct {
	db *Mock
}

type Mock struct {
	Users    map[uuid.UUID]models.User
	Tomatoes map[uuid.UUID]models.Tomato
}

func NewRepository() *Repository {
	db := Mock{}
	db.Users = make(map[uuid.UUID]models.User)
	db.Tomatoes = make(map[uuid.UUID]models.Tomato)

	return &Repository{
		db: &db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	_, ok := r.db.Users[user.Id]
	if !ok {
		return uuid.Nil, errors.New("already exists")
	}
	r.db.Users[user.Id] = user
	return user.Id, nil
}
func (r *Repository) EditUser(ctx context.Context, user models.User) error {
	_, ok := r.db.Users[user.Id]
	if !ok {
		return errors.New("not found")
	}
	r.db.Users[user.Id] = user
	return nil
}
func (r *Repository) GetUser(ctx context.Context, uuid uuid.UUID) (models.User, error) {
	_, ok := r.db.Users[uuid]
	if !ok {
		return models.User{}, errors.New("not found")
	}
	return r.db.Users[uuid], nil
}
func (r *Repository) CreateTomato(ctx context.Context, tomato models.Tomato) (uuid.UUID, error) {
	_, ok := r.db.Users[tomato.Id]
	if !ok {
		return uuid.Nil, errors.New("already exists")
	}
	r.db.Tomatoes[tomato.Id] = tomato
	return tomato.Id, nil
}
func (r *Repository) StartTomato(ctx context.Context, uuid uuid.UUID, start int64) error {
	t, ok := r.db.Tomatoes[uuid]
	if !ok {
		return errors.New("not found")
	}
	t.TimerStart = start
	r.db.Tomatoes[uuid] = t
	return nil
}
func (r *Repository) GetTomato(ctx context.Context, uuid uuid.UUID) (models.Tomato, error) {
	t, ok := r.db.Tomatoes[uuid]
	if !ok {
		return t, errors.New("not found")
	}
	return t, nil
}
func (r *Repository) DeleteTomato(ctx context.Context, uuid uuid.UUID) error {
	_, ok := r.db.Tomatoes[uuid]
	if !ok {
		return errors.New("not found")
	}
	delete(r.db.Tomatoes, uuid)
	return nil
}
func (r *Repository) GetTomatoNltx(ctx context.Context, uuid2 uuid.UUID) (models.TomatoNltx, error) {
	return models.TomatoNltx{}, nil
}
