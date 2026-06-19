package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zhalisher/ip-task-manager/internal/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "email, password and name are required", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 8 {
		http.Error(w, "password must be at least 8 symbols lenght", http.StatusBadRequest)
		return
	}
	user, err := h.authUsecase.Register(r.Context(), req.Email, req.Password, req.Name)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password fields are required", http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, err := h.authUsecase.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
