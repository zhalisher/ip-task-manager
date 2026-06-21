package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
)

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	category.ID = uuid.New()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, `
		INSERT INTO categories (id, user_id, name, color, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, category.ID, category.UserID, category.Name, category.Color, category.CreatedAt, category.UpdatedAt)
	return err
}

func (r *categoryRepository) Update(ctx context.Context, category *model.Category) error {
	category.UpdatedAt = time.Now()
	_, err := r.db.Exec(ctx,
		`UPDATE categories SET name=$1, color=$2, updated_at=$3 WHERE id=$4`,
		category.Name, category.Color, category.UpdatedAt, category.ID,
	)
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`DELETE FROM categories WHERE id=$1`,
		id,
	)
	return err
}

func (r *categoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	category := &model.Category{}

	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, name, color, created_at, updated_at FROM categories WHERE id=$1`,
		id,
	).Scan(&category.ID, &category.UserID, &category.Name, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) GetAll(ctx context.Context, userID uuid.UUID) ([]*model.Category, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, name, color, created_at, updated_at FROM categories WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make([]*model.Category, 0)
	for rows.Next() {
		category := &model.Category{}
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.Color,
			&category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil

}
