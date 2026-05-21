package model

import "time"

type Wallet struct {
	ID                string         `json:"id" db:"id"`
	OrganizationID    string         `json:"organization_id" db:"organization_id"`
	ExternalUserID    string         `json:"external_user_id" db:"external_user_id"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty" db:"external_user_email"`
	Currency          string         `json:"currency" db:"currency"`
	Balance           string         `json:"balance" db:"balance"`
	LedgerBalance     string         `json:"ledger_balance" db:"ledger_balance"`
	LockBalance       string         `json:"lock_balance" db:"lock_balance"`
	IsActive          bool           `json:"is_active" db:"is_active"`
	MetaData          map[string]any `json:"metadata,omitempty" db:"metadata"`
	CreatedAt         time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at" db:"updated_at"`
}
