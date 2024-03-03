// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
}

type Entry struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	AccountID int64     `json:"account_id"`
	// can be negative or positive
	Amount int64 `json:"amount"`
}

type Transfer struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	// must be positive
	Amount int64 `json:"amount"`
}
