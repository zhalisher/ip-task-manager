package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
)

type taskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	task.ID = uuid.New()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx,
		`INSERT INTO tasks (id, user_id, category_id, title, description, status, priority, due_date, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		task.ID, task.UserID, task.CategoryID, task.Title, task.Description, task.Status,
		task.Priority, task.DueDate, task.CreatedAt, task.UpdatedAt,
	)
	return err
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task) error {
	task.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx,
		`UPDATE tasks SET title=$1, description=$2, status=$3, priority=$4, due_date=$5, updated_at=$6 WHERE id=$7`,
		task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.UpdatedAt, task.ID,
	)
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`DELETE FROM tasks WHERE id=$1`,
		id,
	)
	return err
}

func (r *taskRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	task := &model.Task{}

	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, category_id, title, description, status, priority, due_date, created_at, updated_at
		FROM tasks WHERE id=$1`,
		id,
	).Scan(&task.ID, &task.UserID, &task.CategoryID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.DueDate,
		&task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (r *taskRepository) GetAll(ctx context.Context, userID uuid.UUID, filter repository.TaskFilter) ([]*model.Task, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, category_id, title, description, status, priority, due_date, created_at, updated_at
		FROM tasks WHERE user_id=$1 AND status=$2 AND priority=$3 AND title ILIKE $4
		ORDER BY created_at DESC LIMIT $5 OFFSET $6`,
		userID, filter.Status, filter.Priority, "%"+filter.Search+"%", filter.Limit, (filter.Page-1)*filter.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		rows.Scan(&task.ID, &task.UserID, &task.CategoryID, &task.Title, &task.Description, &task.Status, &task.Priority,
			&task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		tasks = append(tasks, task)
	}
	return tasks, nil
}
