package model

type Organization struct {
	ID               string  `json:"id"`
	OrganizationName string  `json:"organization_name"`
	RegisteredName   string  `json:"registered_name"`
	PhoneNumber      string  `json:"phone_number"`
	Email            string  `json:"email"`
	Password         string  `json:"-"`
	APIKey           string  `json:"api_key"`
	APISecret        string  `json:"api_secret"`
	WebsiteURL       *string `json:"website_url,omitempty"`
	WebhookURL       *string `json:"webhook_url,omitempty"`
}
