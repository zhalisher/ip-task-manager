package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
	"github.com/zhalisher/ip-task-manager/internal/usecase"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CategoryID  uuid.UUID `json:"category_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Priority    string    `json:"priority"`
		DueDate     time.Time `json:"due_date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Title == "" || req.Status == "" || req.Priority == "" {
		http.Error(w, "several fields are missing", http.StatusBadRequest)
		return
	}
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	err = h.taskUsecase.Create(r.Context(), &model.Task{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	})
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"message": "task created"})
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Priority    string    `json:"priority"`
		DueDate     time.Time `json:"due_date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = h.taskUsecase.Update(r.Context(), &model.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	})
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "task updated"})
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.taskUsecase.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "task deleted"})
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	task, err := h.taskUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	filter := repository.TaskFilter{
		Search:   r.URL.Query().Get("search"),
		Status:   r.URL.Query().Get("status"),
		Priority: r.URL.Query().Get("priority"),
		Page:     1,
		Limit:    10,
	}
	tasks, err := h.taskUsecase.GetAll(r.Context(), userID, filter)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, tasks)
}
