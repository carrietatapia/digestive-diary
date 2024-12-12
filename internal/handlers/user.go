package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/service/interfaces"

	chi "github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	userCreated, err := h.UserService.CreateUser(context.Background(), &user)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userCreated)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUser(context.Background(), id)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Error getting user from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	userUpdated, err := h.UserService.UpdateUser(context.Background(), id, &user)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Error updating user in database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userUpdated)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	deleted, err := h.UserService.DeleteUser(context.Background(), id)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Error deleting user from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"deleted": deleted})
}
