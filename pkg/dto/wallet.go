package dto

type CreateWalletRequest struct {
	OrganizationID    string         `json:"organization_id" validate:"required"`
	ExternalUserID    string         `json:"external_user_id" validate:"required"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty"`
	Currency          string         `json:"currency" validate:"required"`
	MetaData          map[string]any `json:"metadata,omitempty"`
}

type CreateWalletResponse struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organization_id"`
	ExternalUserID    string         `json:"external_user_id"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty"`
	Currency          string         `json:"currency"`
	Balance           string         `json:"balance"`
	LedgerBalance     string         `json:"ledger_balance"`
	MetaData          map[string]any `json:"metadata,omitempty"`
	IsActive          bool           `json:"is_active"`
	CreatedAt         string         `json:"created_at"`
}
