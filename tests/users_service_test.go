package tests

import (
	"context"
	"testing"

	"github.com/eduardocfalcao/money-tracker/internal/users"
	"github.com/eduardocfalcao/money-tracker/internal/users/models"
	_ "github.com/golang-migrate/migrate/v4/source/file" // used by migrator
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	ctx := context.Background()
	stage := createTestStage()
	defer stage.CleanUp(ctx)
	sut := users.NewService(stage.UsersRepository)

	err := sut.CreateUser(ctx, models.CreateUserRequest{
		Name:            "Test User Name",
		Email:           "user@test.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	})
	require.Nil(t, err)

	users, err := stage.UsersRepository.ListUsers(ctx)

	require.Nil(t, err)
	require.Len(t, users, 1)
	require.Equal(t, "Test User Name", users[0].Name)
	require.Equal(t, "user@test.com", users[0].Email)
}
