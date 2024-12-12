package impl

import (
	"context"
	"errors"

	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/repository/postgres"
)

type UserServiceImpl struct {
	Repo *postgres.UserRepositoryImpl
}

func NewUserServiceImpl(repo *postgres.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Email == "" || user.PasswordHash == "" {
		return nil, errors.New("email and password_hash are required")
	}

	id, err := s.Repo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = id
	user.PasswordHash = "********" // Do not expose sensitive data
	return user, nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserServiceImpl) UpdateUser(ctx context.Context, id string, user *models.User) (*models.User, error) {
	err := s.Repo.Update(ctx, id, user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	user.PasswordHash = "********" // Do not expose sensitive data
	return user, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id string) (bool, error) {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
