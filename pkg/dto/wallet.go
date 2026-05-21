package dto

type CreateWalletRequest struct {
	ExternalUserID    string         `json:"external_user_id" validate:"required"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty"`
	AccountNumber     string         `json:"account_number" validate:"required"`
	AccountName       *string        `json:"account_name,omitempty"`
	Currency          string         `json:"currency" validate:"required"`
	MetaData          map[string]any `json:"metadata,omitempty"`
}
type WalletWithAccountResponse struct {
	Wallet  WalletResponse  `json:"wallet"`
	Account AccountResponse `json:"account"`
}

type WalletResponse struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organization_id"`
	ExternalUserID    string         `json:"external_user_id"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty"`
	Currency          string         `json:"currency"`
	MetaData          map[string]any `json:"metadata,omitempty"`
	IsActive          bool           `json:"is_active"`
	CreatedAt         string         `json:"created_at"`
}

type AccountResponse struct {
	ID               string  `json:"id"`
	WalletID         string  `json:"wallet_id"`
	AccountNumber    string  `json:"account_number"`
	AccountName      *string `json:"account_name,omitempty"`
	Currency         string  `json:"currency"`
	LedgerBalance    string  `json:"ledger_balance"`
	LockedBalance    string  `json:"locked_balance"`
	AvailableBalance string  `json:"available_balance"`
	IsActive         bool    `json:"is_active"`
	CreatedAt        string  `json:"created_at"`
}
