package app

import (
	"context"
	"log"
	"net/http"
	"os"

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
	userUsecase := usecase.NewUserUsecase(userRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	// router
	router := delivery.NewRouter(authHandler, taskHandler, categoryHandler, userHandler, cfg.JWTSecret)

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
