package servis

import (
	"context"
	"github.com/google/uuid"
	"tomato/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
	GetUserByNickOrEmail(ctx context.Context, nick *string, email *string) (*models.User, error)
	CreateUserSession(ctx context.Context, userId uuid.UUID, session uuid.UUID) (uuid.UUID, error)
	DeleteUserSession(ctx context.Context, session uuid.UUID) error
	GetUserBySession(ctx context.Context, session uuid.UUID) (*models.User, error)
	// Habit
	CreateHabit(ctx context.Context, habit models.Habit) (uuid.UUID, error)
	GetHabits(ctx context.Context, userId uuid.UUID, time int64) ([]*models.Habit, error)
	CompletedHabit(ctx context.Context, habitId uuid.UUID) error

	//Task
	CreateTask(ctx context.Context, task models.Task) (uuid.UUID, error)
	GetTasks(ctx context.Context, time int64) ([]*models.Task, error)
	CompletedTask(ctx context.Context, taskId uuid.UUID) error

	// Tomato
	CreateTomato(ctx context.Context, tomato models.Tomato) (uuid.UUID, error)
	GetTomatoes(ctx context.Context, userId uuid.UUID, time int64) ([]*models.Tomato, error)
	StartTomato(ctx context.Context, tomatoId uuid.UUID) error
}
