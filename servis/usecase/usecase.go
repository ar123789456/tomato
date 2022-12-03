package usecase

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"tomato/models"
	"tomato/servis"
)

type UseCase struct {
	repository servis.Repository
}

// NewUseCase UseCase constructor
func NewUseCase(repository servis.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

// User
func (uc *UseCase) CreateUser(user *models.User, ctx context.Context) error {
	user.Id = uuid.New()
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	_, err = uc.repository.CreateUser(ctx, *user)
	return err
}

func (uc *UseCase) SignIn(user *models.User, ctx context.Context) (string, error) {
	user, err := uc.repository.GetUserByNickOrEmail(ctx, &user.Nick, user.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	session := uuid.New()
	_, err = uc.repository.CreateUserSession(ctx, user.Id, session)
	return session.String(), err
}

func (uc *UseCase) SignOut(ctx context.Context) error {
	session, ok := ctx.Value("session").(uuid.UUID)
	if !ok {
		return servis.ErrSessionNotFound
	}
	return uc.repository.DeleteUserSession(ctx, session)
}

func (uc *UseCase) GetUser(ctx context.Context) (*models.User, error) {
	session, ok := ctx.Value("session").(uuid.UUID)
	if !ok {
		return nil, servis.ErrSessionNotFound
	}
	return uc.repository.GetUserBySession(ctx, session)
}

// Habit
func (uc *UseCase) CreateHabit(habit *models.Habit, ctx context.Context) error {
	habit.Id = uuid.New()
	_, err := uc.repository.CreateHabit(ctx, *habit)
	return err
}

func (uc *UseCase) GetHabits(time int64, ctx context.Context) ([]*models.Habit, error) {
	user, err := uc.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return uc.repository.GetHabits(ctx, user.Id, time)
}

func (uc *UseCase) CompletedHabit(habitId string, ctx context.Context) error {
	panic("implement me")
}

// Task
func (uc *UseCase) CreateTask(task *models.Task, ctx context.Context) error {
	panic("implement me")
}

func (uc *UseCase) GetTasks(session string, time int64, ctx context.Context) ([]*models.Task, error) {
	panic("implement me")
}

func (uc *UseCase) CompletedTask(taskId string, ctx context.Context) error {
	panic("implement me")
}

// Tomato
func (uc *UseCase) CreateTomato(tomato *models.Tomato, ctx context.Context) error {
	panic("implement me")
}

func (uc *UseCase) GetTomatoes(time int64, ctx context.Context) ([]*models.Tomato, error) {
	panic("implement me")
}

func (uc *UseCase) StartTomato(tomatoId string, ctx context.Context) error {
	panic("implement me")
}
