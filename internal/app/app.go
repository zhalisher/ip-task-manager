package app

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zhalisher/ip-task-manager/config"
	delivery "github.com/zhalisher/ip-task-manager/internal/delivery/http"
	handler "github.com/zhalisher/ip-task-manager/internal/delivery/http/handler"
	"github.com/zhalisher/ip-task-manager/internal/infrastructure/db/postgres"
	"github.com/zhalisher/ip-task-manager/internal/usecase"
)

func Run(cfg *config.Config) {
	// db
	pool, err := pgxpool.New(context.Background(), cfg.PostgresURI)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	// repositories
	userRepo := postgres.NewUserRepository(pool)
	taskRepo := postgres.NewTaskRepository(pool)
	categoryRepo := postgres.NewCategoryRepository(pool)

	// usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret, cfg.JWTAccessExpMin)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	// router
	router := delivery.NewRouter(authHandler, taskHandler, categoryHandler, cfg.JWTSecret)

	// server
	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
