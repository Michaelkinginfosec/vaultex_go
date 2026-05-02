package repository

import (
	"context"
	"fmt"
	"vaultex/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	UpdateUserById(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUserById(ctx context.Context, id string) error
}

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID)

}
func (r *repo) FindByID(ctx context.Context, id string) (*model.User, error) {

	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repo) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email FROM users WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {

		return nil, err
	}
	return &user, nil
}

func (r *repo) UpdateUserById(ctx context.Context, id string, user *model.User) (*model.User, error) {
	var updatedUser model.User

	query := `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3
		RETURNING id, name, email
	`

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, id).
		Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &updatedUser, nil

}

func (r *repo) DeleteUserById(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("User with id %s not found", id)
	}
	return nil
}
