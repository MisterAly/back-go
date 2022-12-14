// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID       int64 `json:"id"`
	PersonID int64 `json:"personID"`
	// can be negative
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Expense struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"accountID"`
	Value     int64     `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

type Person struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Document  string         `json:"document"`
	Phone     sql.NullString `json:"phone"`
	CreatedAt time.Time      `json:"createdAt"`
}
