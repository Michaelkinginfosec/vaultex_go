package dto

type CreateWalletRequest struct {
	OrganizationID    string         `json:"organization_id" validate:"required"`
	ExternalUserID    string         `json:"external_user_id" validate:"required"`
	ExternalUserEmail *string        `json:"external_user_email,omitempty"`
	Currency          string         `json:"currency" validate:"required"`
	MetaData          map[string]any `json:"metadata,omitempty"`
}
