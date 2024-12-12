package interfaces

import (
	"context"

	"github.com/carrietatapia/digestivediary/internal/domain/models"
)

type UserRepository interface {
	Insert(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}
