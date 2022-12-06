package usecase

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	return uc.repository.CompletedHabit(ctx, uuid.MustParse(habitId))
}

// Task
func (uc *UseCase) CreateTask(task *models.Task, ctx context.Context) error {
	task.CreatedAt = time.Now().Unix()
	task.Id = uuid.New()
	user, err := uc.GetUser(ctx)
	if err != nil {
		return err
	}
	task.UserId = user.Id
	_, err = uc.repository.CreateTask(ctx, *task)
	return err
}

func (uc *UseCase) GetTasks(time int64, ctx context.Context) ([]*models.Task, error) {
	return uc.repository.GetTasks(ctx, time)
}

func (uc *UseCase) CompletedTask(taskId string, ctx context.Context) error {
	return uc.repository.CompletedTask(ctx, uuid.MustParse(taskId))
}

// Tomato
func (uc *UseCase) CreateTomato(tomato *models.Tomato, ctx context.Context) error {
	user, err := uc.GetUser(ctx)
	if err != nil {
		return err
	}
	tomato.Id = uuid.New()
	tomato.UserId = user.Id
	tomato.CreateTime = time.Now().Unix()
	_, err = uc.repository.CreateTomato(ctx, *tomato)
	return err
}

func (uc *UseCase) GetTomatoes(time int64, ctx context.Context) ([]*models.Tomato, error) {
	user, err := uc.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return uc.repository.GetTomatoes(ctx, user.Id, time)
}

func (uc *UseCase) StartTomato(tomatoId string, ctx context.Context) error {
	return uc.repository.StartTomato(ctx, uuid.MustParse(tomatoId))
}
