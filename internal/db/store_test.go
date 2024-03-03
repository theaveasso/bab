package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"theaveasso.bab/internal/utility"
)

func TestTransferTX(t *testing.T) {
	testStore := NewStore(testDB)

	sender, _ := createRandomAccount(t)
	receiver, _ := createRandomAccount(t)

	n := 5

	errs := make(chan error)
	results := make(chan TransferTXResult)

	var amount int64
	for i := 0; i < n; i++ {
		amount = utility.RandomInt(5, 10)
		go func() {
			arg := TransferTXParams{
				FromAccountID: sender.ID,
				ToAccountID:   receiver.ID,
				Amount:        amount,
			}
			result, err := testStore.TransferTX(context.Background(), arg)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, sender.ID, transfer.FromAccountID)
		require.Equal(t, receiver.ID, transfer.ToAccountID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, sender.ID, fromEntry.AccountID)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.ID)

		_, err = testStore.GetEntry(context.Background(), sender.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, receiver.ID, toEntry.AccountID)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.ID)

		_, err = testStore.GetEntry(context.Background(), receiver.ID)
		require.NoError(t, err)
	}
}
