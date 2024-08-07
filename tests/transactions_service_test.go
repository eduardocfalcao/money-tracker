package tests

import (
	"context"
	"os"
	"testing"

	"github.com/eduardocfalcao/money-tracker/internal/api"
	"github.com/eduardocfalcao/money-tracker/internal/transactions"
	"github.com/eduardocfalcao/money-tracker/internal/transactions/repository"
	"github.com/eduardocfalcao/money-tracker/internal/users/models"
	"github.com/stretchr/testify/require"
)

func Test_ImportOFXFile(t *testing.T) {
	ctx := context.Background()
	stage := createTestStage()
	defer stage.CleanUp(ctx)

	sut := transactions.NewService(stage.TransactionsRepository)

	// load file
	file, err := os.Open("testdata/test-1.ofx")
	require.Nil(t, err)
	defer file.Close()

	ctx = api.SetContextUser(ctx, &api.ApiUser{
		UserID: 1,
	})

	stage.UsersRepository.CreateUser(ctx, models.User{
		Name:         "test_user",
		Email:        "test@mail.com",
		Passwordhash: "1234",
	})

	err = sut.ImportOFXFile(ctx, file)
	require.Nil(t, err)

	transactions, err := stage.TransactionsRepository.GetTransactionsByUser(ctx, repository.GetTransactionsByUserParams{
		PageParams: repository.PageParams{
			Limit:  1000,
			Offset: 0,
		},
		UserID: 1,
	})

	require.Nil(t, err, err)
	require.Len(t, transactions, 15)
}
