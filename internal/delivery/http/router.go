package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/zhalisher/ip-task-manager/internal/delivery/http/handler"
	"github.com/zhalisher/ip-task-manager/internal/middleware"
)

func NewRouter(
	authHandler *handler.AuthHandler,
	taskHandler *handler.TaskHandler,
	categoryHandler *handler.CategoryHandler,
	jwtSecret string,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recovery)

	// public routes
	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)

	// secuired routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(jwtSecret))

		r.Post("/tasks", taskHandler.Create)
		r.Get("/tasks", taskHandler.GetAll)
		r.Get("/tasks/{id}", taskHandler.GetByID)
		r.Put("/tasks/{id}", taskHandler.Update)
		r.Delete("/tasks/{id}", taskHandler.Delete)

		r.Post("/categories", categoryHandler.Create)
		r.Get("/categories", categoryHandler.GetAll)
		r.Get("/categories/{id}", categoryHandler.GetByID)
		r.Put("/categories/{id}", categoryHandler.Update)
		r.Delete("/categories/{id}", categoryHandler.Delete)
	})
	return r
}
