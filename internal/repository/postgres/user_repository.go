package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Insert(ctx context.Context, user *models.User) (string, error) {
	log.Println("Before inserting user:", user)
	log.Println("ctx:", ctx)
	log.Println("db:", r.db)

	query := `
        INSERT INTO users (id, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()

	err := db.DB.QueryRow(ctx, query, id, user.Email, user.PasswordHash, now).Scan(&id)
	if err != nil {
		log.Println("Error inserting user:", err)
		return "", err
	}
	log.Println("User inserted successfully with ID:", id)
	return id.String(), nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, email, created_at FROM users WHERE id = $1;`
	var user models.User
	err := db.DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, id string, user *models.User) error {
	fields := utils.BuildUpdateMap(user)
	log.Println("fields:", fields)

	return utils.ExecuteUpdate(ctx, r.db, "users", id, fields)
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM users WHERE id = $1;
	`
	_, err := db.DB.Exec(ctx, query, id)
	return err
}
