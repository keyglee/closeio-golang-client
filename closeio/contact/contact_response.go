package contact

import customfield "github.com/keyglee/closeio-golang-client/closeio/customField"

type ContactResponse struct {
	Name           string                    `json:"name"`
	Title          string                    `json:"title"`
	DateUpdated    string                    `json:"date_updated"`
	Phones         []PhoneResponse           `json:"phones"`
	CustomFields   []customfield.CustomField `json:"custom_fields"`
	CreatedBy      string                    `json:"created_by"`
	ID             string                    `json:"id"`
	OrganizationId string                    `json:"organization_id"`
	DateCreated    string                    `json:"date_created"`
	Emails         []EmailResponse           `json:"emails"`
	UpdatedBy      string                    `json:"updated_by"`
}

type EmailResponse struct {
	Type           string `json:"type"`
	Email          string `json:"email"`
	IsUnsubscribed bool   `json:"is_unsubscribed"`
}

type PhoneResponse struct {
	Type           bool   `json:"type"`
	Phone          string `json:"phone"`
	PhoneFormatted string `json:"phone_formatted"`
}
