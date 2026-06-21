package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/usecase"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Color == "" {
		http.Error(w, "name and color are required", http.StatusBadRequest)
		return
	}
	err := h.categoryUsecase.Create(r.Context(), &model.Category{
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"message": "category created"})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Color == "" {
		http.Error(w, "name and color are required", http.StatusBadRequest)
		return
	}
	err := h.categoryUsecase.Update(r.Context(), &model.Category{
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "category updated"})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.categoryUsecase.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "category deleted"})
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	category, err := h.categoryUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "category not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	categories, err := h.categoryUsecase.GetAll(r.Context(), userID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, categories)
}
