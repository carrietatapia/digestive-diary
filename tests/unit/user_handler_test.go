package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/handlers"
	"github.com/carrietatapia/digestivediary/internal/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockUserService)

	newUser := models.User{
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
	}

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
		Return(&models.User{ID: uuid.New().String(), Email: newUser.Email}, nil)

	payload, _ := json.Marshal(newUser)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	handler.CreateUserHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var userResponse models.User
	json.NewDecoder(res.Body).Decode(&userResponse)

	assert.NotEmpty(t, userResponse.ID)
	assert.Equal(t, newUser.Email, userResponse.Email)
}

func TestGetUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockUserService)

	userID := uuid.New().String()
	mockUser := &models.User{ID: userID, Email: "test@example.com"}

	mockUserService.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", userID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

	w := httptest.NewRecorder()

	handler.GetUserHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var userResponse models.User
	err := json.NewDecoder(res.Body).Decode(&userResponse)
	assert.NoError(t, err, "Failed to decode response body")

	assert.Equal(t, mockUser.ID, userResponse.ID)
	assert.Equal(t, mockUser.Email, userResponse.Email)
}
