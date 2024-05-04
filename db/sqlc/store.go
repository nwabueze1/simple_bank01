package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) executeTransaction(ctx context.Context, callback func(queries *Queries) error) error {
	transaction, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	transactionQueries := New(transaction)

	err = callback(transactionQueries)
	if err != nil {
		rollbackError := transaction.Rollback()
		if rollbackError != nil {
			return fmt.Errorf("rollbackError: %v, error:%v", rollbackError, err)
		}

		return err
	}

	return transaction.Commit()
}

type TransferTransactionParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTransactionResult struct {
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	Transfer    Transfer `json:"transfer"`
}

func (store *Store) TransferTx(ctx context.Context, args TransferTransactionParams) (TransferTransactionResult, error) {
	var result TransferTransactionResult

	err := store.executeTransaction(ctx, func(queries *Queries) error {
		var err error

		//create transfer
		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			Amount:        args.Amount,
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
		})

		if err != nil {
			return err
		}

		//create from account entry
		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			Amount:    -args.Amount,
			AccountID: args.FromAccountID,
		})

		if err != nil {
			return err
		}

		//create to account entry
		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			Amount:    args.Amount,
			AccountID: args.ToAccountID,
		})
		if err != nil {
			return err
		}

		if args.FromAccountID < args.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, args.FromAccountID,
				-args.Amount, args.ToAccountID, args.Amount,
			)

			//TODO: UPDATE THE FROM BALANCE AND TO BALANCE
			return nil
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, queries, args.ToAccountID,
				args.Amount, args.FromAccountID, -args.Amount,
			)
		}
		return nil
	})
	return result, err
}

func addMoney(ctx context.Context,
	q *Queries,
	account1ID int64,
	amount1 int64,
	account2ID int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount1,
		ID:     account1ID,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount2,
		ID:     account2ID,
	})

	if err != nil {
		return
	}

	return
}
