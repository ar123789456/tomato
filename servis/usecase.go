package servis

import (
	"context"
	"errors"
	"tomato/models"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type UseCase interface {
	// User
	CreateUser(user *models.User, ctx context.Context) error
	SignIn(user *models.User, ctx context.Context) (string, error)
	SignOut(ctx context.Context) error
	GetUser(ctx context.Context) (*models.User, error)

	// Habit
	CreateHabit(habit *models.Habit, ctx context.Context) error
	GetHabits(time int64, ctx context.Context) ([]*models.Habit, error)
	CompletedHabit(habitId string, ctx context.Context) error

	// Task
	CreateTask(task *models.Task, ctx context.Context) error
	GetTasks(time int64, ctx context.Context) ([]*models.Task, error)
	CompletedTask(taskId string, ctx context.Context) error

	// Tomato
	CreateTomato(tomato *models.Tomato, ctx context.Context) error
	GetTomatoes(time int64, ctx context.Context) ([]*models.Tomato, error)
	StartTomato(tomatoId string, ctx context.Context) error
}
