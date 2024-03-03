// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING id, created_at, from_account_id, to_account_id, amount
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.queryRow(ctx, q.createTransferStmt, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, created_at, from_account_id, to_account_id, amount FROM transfers 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.queryRow(ctx, q.getTransferStmt, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
	)
	return i, err
}
