package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"theaveasso.bab/internal/utility"
)

func createRandomAccount(t *testing.T) (Account, error) {
    userCreated := createRandomUser(t)
	args := CreateAccountParams{
		Username: userCreated.Username,
		Balance:  utility.RandomBalance(),
		Currency: utility.RandomCurrency(),
	}

    account, err :=  testQueries.CreateAccount(context.TODO(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Username, account.Username)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

    return account, nil
}

func TestCreateAccount(t *testing.T) {
    createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
    account1, _ := createRandomAccount(t)
    account2, err := testQueries.GetAccount(context.TODO(), account1.ID)

    require.NoError(t, err)
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Username, account2.Username)
    require.Equal(t, account1.Currency, account2.Currency)
    require.Equal(t, account1.Balance, account2.Balance)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
    account1, _ := createRandomAccount(t)

	args := UpdateAccountBalanceParams{
		Balance:  utility.RandomBalance(),
        ID:       account1.ID,
	}

    account2, err := testQueries.UpdateAccountBalance(context.TODO(), args)

    require.NoError(t, err)
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Username, account2.Username)
    require.Equal(t, account1.Currency, account2.Currency)
    require.Equal(t, args.Balance, account2.Balance)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func  TestDeleteAccount(t *testing.T) {
    account1, _ := createRandomAccount(t)

    err := testQueries.DeleteAccount(context.TODO(), account1.ID)
    require.NoError(t, err)

    account2, err := testQueries.GetAccount(context.TODO(), account1.ID)
    require.Error(t, err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
    require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
    for i := 0; i < 10; i++ {
        createRandomAccount(t)
    }

    accounts, err := testQueries.ListAccounts(context.TODO(), ListAccountsParams{
        Limit: 5,
        Offset: 5,
    })

    require.NoError(t, err)
    require.Equal(t, 5, len(accounts))
    for _, account := range accounts {
        require.NotEmpty(t, account)
    }
}
