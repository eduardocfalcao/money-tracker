package tests

import (
	"context"
	"testing"

	"github.com/eduardocfalcao/money-tracker/database/queries"
	_ "github.com/golang-migrate/migrate/v4/source/file" // used by migrator
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	ctx := context.Background()
	stage := createTestStage()
	defer stage.CleanUp(ctx)

	err := stage.Repository.CreateUser(ctx, queries.CreateUserParams{
		Name:         "Test User Name",
		Email:        "user@test.com",
		Passwordhash: "aaaaaa",
		Salt:         "123",
	})
	require.Nil(t, err)
}
