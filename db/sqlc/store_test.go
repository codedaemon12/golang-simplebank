package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	accounts1 := CreateRandomAccount(t)
	accounts2 := CreateRandomAccount(t)
	fmt.Printf("FromAccount::%v toAccount::%v\n", accounts1.ID, accounts2.ID)
	fmt.Println(">>> before", accounts1.Balance, accounts2.Balance)
	n := 5

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: accounts1.ID,
				ToAccountID:   accounts2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer

		transfer := result.Transfer

		require.NotEmpty(t, transfer)
		require.Equal(t, accounts1.ID, transfer.FromAccountID)
		require.Equal(t, accounts2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTranfer(context.Background(), transfer.FromAccountID)
		require.NoError(t, err)

		// // check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.AccountID, accounts1.ID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.AccountID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, accounts2.ID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.AccountID)
		require.NoError(t, err)

		// check accounts

		fromAccount := result.FromAccount
		fmt.Println("fromAccount", fromAccount.ID)
		require.NotEmpty(t, fromAccount)
		require.Equal(t, accounts1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		fmt.Println("toAccount", toAccount.ID)
		require.NotEmpty(t, toAccount)
		require.Equal(t, accounts2.ID, toAccount.ID)

		//check balance
		fmt.Println(">>> tx", fromAccount.Balance, toAccount.Balance)
		diff1 := accounts1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - accounts2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err1 := testQueries.GetAccount(context.Background(), accounts1.ID)
	require.NoError(t, err1)
	updatedAccount2, err2 := testQueries.GetAccount(context.Background(), accounts2.ID)
	require.NoError(t, err2)

	require.Equal(t, accounts1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, accounts2.Balance+int64(n)*amount, updatedAccount2.Balance)

	fmt.Println(">>> after", updatedAccount1.Balance, updatedAccount2.Balance)
}
