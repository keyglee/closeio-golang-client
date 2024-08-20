package contact

import customfield "github.com/keyglee/closeio-golang-client/closeio/customField"

type ContactResponse struct {
	Name            string                     `json:"name"`
	Title           string                     `json:"title"`
	Date_updated    string                     `json:"date_updated"`
	Phones          []PhoneResponse            `json:"phones"`
	Custom_fields   []customfield.Custom_Field `json:"custom_fields"`
	Created_by      string                     `json:"created_by"`
	Id              string                     `json:"id"`
	Organization_id string                     `json:"organization_id"`
	Date_created    string                     `json:"date_created"`
	Emails          []EmailResponse            `json:"emails"`
	Updated_by      string                     `json:"updated_by"`
}

type EmailResponse struct {
	Type            string `json:"type"`
	Email           string `json:"email"`
	Is_unsubscribed bool   `json:"is_unsubscribed"`
}

type PhoneResponse struct {
	Type            bool   `json:"type"`
	Phone           string `json:"phone"`
	Phone_formatted string `json:"phone_formatted"`
}
