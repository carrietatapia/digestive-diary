package repository_test

import (
	"context"
	"testing"

	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/domain/models"
	"github.com/carrietatapia/digestivediary/internal/repository/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	testDB   *pgxpool.Pool
	userRepo *postgres.UserRepositoryImpl
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	var err error
	suite.testDB, err = db.NewTestDB()
	require.NoError(suite.T(), err)
	db.DB = suite.testDB
	suite.userRepo = postgres.NewUserRepository(suite.testDB)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	suite.testDB.Close()
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db.CleanupDB(suite.testDB)
}

func (suite *UserRepositoryTestSuite) TestInsertFindByID() {
	user := &models.User{
		Email:        "integration2@example.com",
		PasswordHash: "hashed_password",
	}

	ctx := context.Background()
	userID, err := suite.userRepo.Insert(ctx, user)
	require.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), userID)

	foundUser, err := suite.userRepo.FindByID(ctx, userID)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepositoryTestSuite) TestUpdateDelete() {
	newUser := &models.User{
		Email:        "update_test@example.com",
		PasswordHash: "hashed_password",
	}

	ctx := context.Background()
	userID, err := suite.userRepo.Insert(ctx, newUser)
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), userID)

	updatedEmail := "updated_email@example.com"
	newUser.Email = updatedEmail

	err = suite.userRepo.Update(ctx, userID, newUser)
	require.NoError(suite.T(), err)

	updatedUser, err := suite.userRepo.FindByID(ctx, userID)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedEmail, updatedUser.Email)

	err = suite.userRepo.Delete(ctx, userID)
	require.NoError(suite.T(), err)

	deletedUser, err := suite.userRepo.FindByID(ctx, userID)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), deletedUser)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &UserRepositoryTestSuite{})
}
