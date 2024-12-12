package e2e_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/handlers"
	"github.com/carrietatapia/digestivediary/internal/repository/postgres"
	"github.com/carrietatapia/digestivediary/internal/service/impl"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	testDB *pgxpool.Pool
)

func TestMain(m *testing.M) {
	// Setup
	var err error
	testDB, err = db.NewTestDB()
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}

	// Set the global DB variable to the testDB
	db.DB = testDB

	// Run tests
	code := m.Run()

	// Cleanup
	db.CleanupDB(testDB)
	testDB.Close()

	os.Exit(code)
}

func setupRouter() chi.Router {
	// Clean the DB before each test
	db.CleanupDB(testDB)

	userRepo := postgres.NewUserRepository(testDB)
	userService := impl.NewUserServiceImpl(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Post("/users", userHandler.CreateUserHandler)
	r.Get("/users/{id}", userHandler.GetUserHandler)

	return r
}

func TestCreateAndGetUser(t *testing.T) {
	// Setup
	r := setupRouter()

	t.Run("Create User", func(t *testing.T) {
		// Prepare request
		newUser := models.User{
			Email:        "test2@example.com",
			PasswordHash: "hashed_password",
		}
		payload, err := json.Marshal(newUser)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Execute request
		r.ServeHTTP(w, req)

		// Assert response
		res := w.Result()
		defer res.Body.Close()

		require.Equal(t, http.StatusOK, res.StatusCode)
		require.Equal(t, "application/json", res.Header.Get("Content-Type"))

		var createdUser models.User
		err = json.NewDecoder(res.Body).Decode(&createdUser)
		require.NoError(t, err)
		require.NotEmpty(t, createdUser.ID)
		require.Equal(t, newUser.Email, createdUser.Email)

		// Test Get User
		t.Run("Get Created User", func(t *testing.T) {
			req = httptest.NewRequest(http.MethodGet, "/users/"+createdUser.ID, nil)
			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)

			res = w.Result()
			defer res.Body.Close()

			require.Equal(t, http.StatusOK, res.StatusCode)
			require.Equal(t, "application/json", res.Header.Get("Content-Type"))

			var fetchedUser models.User
			err = json.NewDecoder(res.Body).Decode(&fetchedUser)
			require.NoError(t, err)
			require.Equal(t, createdUser.ID, fetchedUser.ID)
			require.Equal(t, createdUser.Email, fetchedUser.Email)
		})
	})
}
