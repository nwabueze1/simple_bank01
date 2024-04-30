package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccoun(t *testing.T) {
	args := CreateAccountParams{
		Owner: "Tom",
		Balance: 100,
		Currency: "NGN",
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t,err)
	require.NotEmpty(t,createdAccount)
	require.Equal(t, args.Owner, createdAccount.Owner)
	require.Equal(t, args.Balance, createdAccount.Balance)
	require.Equal(t, args.Currency, createdAccount.Currency)
}