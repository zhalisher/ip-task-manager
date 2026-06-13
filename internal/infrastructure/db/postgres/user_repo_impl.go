package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx,
		`INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6)`,
		user.ID, user.Email, user.Password, user.Name, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(ctx,
		`UPDATE users SET name=$1, email=$2, updated_at=$3 WHERE id=$4`,
		user.Name, user.Email, user.UpdatedAt, user.ID,
	)
	return err
}
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`DELETE FROM users WHERE id=$1`,
		id,
	)
	return err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(ctx,
		`SELECT id, email, password, name, created_at, updated_at FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(ctx,
		`SELECT id, email, password, name, created_at, updated_at FROM users WHERE id=$1`,
		id,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
