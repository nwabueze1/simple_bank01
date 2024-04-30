package db

import (
	"context"
	"testing"
	"time"

	"fidelis.com/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RadomMoney(),
		Currency: util.RandomCurrency(),
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t,err)
	require.NotEmpty(t,createdAccount)
	require.Equal(t, args.Owner, createdAccount.Owner)
	require.Equal(t, args.Balance, createdAccount.Balance)
	require.Equal(t, args.Currency, createdAccount.Currency)

	return createdAccount
}

func TestCreateAccoun(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T){
	account1 := createTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T){
	//create a random test accounts of 5 params
	for i:=0; i< 5; i++{
		createTestAccount(t)
	}
	args:= ListAccountsParams{
		Limit: 5,
		Offset: 1,
	}

	accounts,err := testQueries.ListAccounts(context.Background(),args )

	require.NoError(t, err)
	require.NotEmpty(t,accounts)
	require.Len(t, accounts, 5)
}

func TestUpdateAccount(t *testing.T){
	account1 := createTestAccount(t)
	args := UpdateAccountParams{
		ID: account1.ID,
		Balance: 600,
	}

	account2, err := testQueries.UpdateAccount(context.Background(),args)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, args.Balance)
}

func TestDeleteAccount(t *testing.T){
	account1 := createTestAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}