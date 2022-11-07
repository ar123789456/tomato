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
	query := "INSERT into users (id, name, secondName, nick, email, photo, class, session) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := r.db.QueryContext(ctx, query, user.Id.String(), user.Name, user.SecondName, user.Nick, user.Email,
		user.Photo, user.Class, user.Session.String())
	return user.Id, err
}
func (r *Repository) EditUser(ctx context.Context, user models.User) error {
	query := "update users set name=$1, secondName=$2, nick=$3, email=$4, photo=$5, class=$6, session=$7 where id=$8"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.SecondName, user.Nick, user.Email,
		user.Photo, user.Class, user.Session.String(), user.Id.String())
	return err
}
func (r *Repository) GetUser(ctx context.Context, uuidUser uuid.UUID) (models.User, error) {
	query := "select * from users where id=$1 or session=$2"
	row := r.db.QueryRowContext(ctx, query, uuidUser.String())
	var user models.User
	var id string
	var session string
	err := row.Scan(&id, &user.Name, &user.SecondName, &user.Nick, &user.Email, &user.Photo, &user.Class, &session)
	if err != nil {
		return user, err
	}
	user.Id, err = uuid.Parse(id)
	if err != nil {
		return user, err
	}
	ses, err := uuid.Parse(session)
	if err != nil {
		return user, err
	}
	user.Session = &ses
	return models.User{}, nil
}
func (r *Repository) CreateTomato(ctx context.Context, tomato models.Tomato) (uuid.UUID, error) {
	query := `begin transaction;
	insert into tomatoes (id, timeStart, createTime, title, context, user_id) 
	values ($1, $2, $3, $4, $5, $6);  
	insert into tags (tag)
`
	r.db.ExecContext(ctx, query, tomato)
	return uuid.Nil, nil
}
func (r *Repository) StartTomato(ctx context.Context, uuid uuid.UUID, start int64) error {
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
