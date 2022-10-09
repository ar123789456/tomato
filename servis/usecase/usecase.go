package usecase

import (
	"context"
	"github.com/google/uuid"
	"strings"
	"time"
	"tomato/models"
	"tomato/servis"
)

type UseCase struct {
	repository servis.Repository
}

func NewUseCase(repository servis.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateUser(ctx context.Context, userIn models.CreateUser) (uuid.UUID, error) {
	var user models.User
	if userIn.SecondName != nil && strings.TrimSpace(*userIn.SecondName) == "" {
		userIn.SecondName = nil
	}
	if userIn.Email != nil && strings.TrimSpace(*userIn.Email) == "" {
		userIn.Email = nil
	}
	id := uuid.New()
	user.Id = id
	user.Name = userIn.Name
	user.SecondName = userIn.SecondName
	user.Nick = userIn.Nick
	user.Class = userIn.Class
	user.Email = userIn.Email
	user.Photo = userIn.Photo

	return uc.repository.CreateUser(ctx, user)
}

func (uc *UseCase) EditUser(ctx context.Context, userIn models.EditUser) error {
	var user models.User
	if userIn.SecondName != nil && strings.TrimSpace(*userIn.SecondName) == "" {
		userIn.SecondName = nil
	}
	if userIn.Email != nil && strings.TrimSpace(*userIn.Email) == "" {
		userIn.Email = nil
	}
	user.Id = userIn.Id
	user.Name = userIn.Name
	user.SecondName = userIn.SecondName
	user.Nick = userIn.Nick
	user.Class = userIn.Class
	user.Email = userIn.Email
	user.Photo = userIn.Photo
	err := user.Validate()
	if err != nil {
		return err
	}
	return uc.repository.EditUser(ctx, user)
}

func (uc *UseCase) GetUser(ctx context.Context, uuid uuid.UUID) (models.User, error) {
	return uc.repository.GetUser(ctx, uuid)
}

func (uc *UseCase) CreateTomato(ctx context.Context, tomato models.CreateTomatoIn) (uuid.UUID, error) {
	var tmt models.Tomato
	tmt.Id = uuid.New()
	tmt.Title = tomato.Title
	tomato.TimerTomato.SetRest()
	tmt.TimerTomato = tomato.TimerTomato
	tmt.CreateTime = time.Now().Unix()
	tmt.Tags = tomato.Tags
	tmt.Context = tomato.Context

	err := tmt.Validate()
	if err != nil {
		return uuid.Nil, err
	}
	return uc.repository.CreateTomato(ctx, tmt)
}
func (uc *UseCase) StartTomato(ctx context.Context, uuid uuid.UUID) error {

	return uc.repository.StartTomato(ctx, uuid, time.Now().Unix())
}
func (uc *UseCase) GetTomato(ctx context.Context, uuid uuid.UUID) (models.Tomato, error) {
	return uc.repository.GetTomato(ctx, uuid)
}
func (uc *UseCase) DeleteTomato(ctx context.Context, uuid uuid.UUID) error {
	return uc.repository.DeleteTomato(ctx, uuid)
}
func (uc *UseCase) GetTomatoNltx(ctx context.Context, uuid uuid.UUID) (models.TomatoNltx, error) {
	return uc.repository.GetTomatoNltx(ctx, uuid)
}
