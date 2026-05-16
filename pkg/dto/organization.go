package dto

type CreateOrganizationRequest struct {
	OrganizationName string `json:"organization_name" binding:"required"`
	RegisteredName   string `json:"registered_name" binding:"required"`
	PhoneNumber      string `json:"phone_number" binding:"required"`
	WebsiteURL       string `json:"website_url"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=8"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateOrganizationRequest struct {
	OrganizationName *string `json:"organization_name"`
	RegisteredName   *string `json:"registered_name"`
	PhoneNumber      *string `json:"phone_number"`
	WebsiteURL       *string `json:"website_url"`
	WebhookURL       *string `json:"webhook_url"`
}

type OrganizationSignupResponse struct {
	ID               string `json:"id"`
	OrganizationName string `json:"organization_name"`
	RegisteredName   string `json:"registered_name"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	APIKey           string `json:"api_key"`
	APISecret        string `json:"api_secret"`
}

type OrganizationLoginResponse struct {
	ID               string `json:"id"`
	OrganizationName string `json:"organization_name"`
	RegisteredName   string `json:"registered_name"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	APIKey           string `json:"api_key"`
}
