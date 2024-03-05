package db

import (
	"context"
	"database/sql"
	"fmt"
)

type SQLStore struct {
	*Queries
	db *sql.DB
}

type Store interface {
    Querier
    TransferTX(ctx context.Context, args TransferTXParams) (TransferTXResult, error)
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes a function within a database transaction
func (s *SQLStore) execTx(ctx context.Context, f func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = f(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %+v, rollback error %+v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

type TransferTXParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferTXResult struct {
	Transfer     Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *SQLStore) TransferTX(ctx context.Context, args TransferTXParams) (TransferTXResult, error) {
	var result TransferTXResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = s.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = s.CreateEntry(ctx, CreateEntryParams{
			Amount:    -args.Amount,
			AccountID: args.FromAccountID,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = s.CreateEntry(ctx, CreateEntryParams{
			Amount:    args.Amount,
			AccountID: args.ToAccountID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
