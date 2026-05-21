package model

import (
	"time"
)

type Wallet struct {
	ID                string         `json:"id" db:"id"`
	OrganizationID    string         `json:"organization_id" db:"organization_id"`
	ExternalUserID    string         `json:"external_user_id" db:"external_user_id"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty" db:"external_user_email"`
	Currency          string         `json:"currency" db:"currency"`
	IsActive          bool           `json:"is_active" db:"is_active"`
	MetaData          map[string]any `json:"metadata,omitempty" db:"metadata"`
	CreatedAt         time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at" db:"updated_at"`
}

type Account struct {
	ID             string    `json:"id" db:"id"`
	WalletID       string    `json:"wallet_id" db:"wallet_id"`
	OrganizationID string    `json:"organization_id" db:"organization_id"`
	AccountNumber  string    `json:"account_number" db:"account_number"`
	AccountName    *string   `json:"account_name,omitempty" db:"account_name"`
	Currency       string    `json:"currency" db:"currency"`
	LedgerBalance  string    `json:"ledger_balance" db:"ledger_balance"`
	LockedBalance  string    `json:"locked_balance" db:"locked_balance"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
type WalletWithAccount struct {
	Wallet  *Wallet  `json:"wallet"`
	Account *Account `json:"account"`
}
