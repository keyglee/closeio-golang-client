package lead

import (
	"github.com/keyglee/closeio-golang-client/closeio/contact"
	customfield "github.com/keyglee/closeio-golang-client/closeio/customField"
)

type Lead struct {
	Name          string                     `json:"name,omitempty"`
	Url           string                     `json:"url,omitempty"`
	Description   string                     `json:"description,omitempty"`
	Status_id     string                     `json:"status_id,omitempty"`
	Contacts      []contact.Contact          `json:"contacts,omitempty"`
	Custom_fields []customfield.Custom_Field `json:"custom_fields,omitempty"`
	Addresses     []Address                  `json:"addresses,omitempty"`
}

type Address struct {
	Label    string `json:"label,omitempty"`
	Address1 string `json:"address_1,omitempty"`
	Address2 string `json:"address_2,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Zipcode  string `json:"zipcode,omitempty"`
	Country  string `json:"country,omitempty"`
}
