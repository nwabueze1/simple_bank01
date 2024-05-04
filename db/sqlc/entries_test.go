package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTestEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}

	entry, err := testQuery.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createTestEntry(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry1 := createTestEntry(t)

	entry2, err := testQuery.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry2.ID, entry1.ID)
	require.Equal(t, entry2.AccountID, entry1.AccountID)
	require.Equal(t, entry2.Amount, entry1.Amount)
	require.WithinDuration(t, entry2.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestEntry(t)
	}

	args := ListEntriesParams{
		Offset: 1,
		Limit:  5,
	}

	entries, err := testQuery.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.CreatedAt)
	}
}
