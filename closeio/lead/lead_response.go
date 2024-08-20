package lead

import (
	"github.com/keyglee/closeio-golang-client/closeio/contact"
	customfield "github.com/keyglee/closeio-golang-client/closeio/customField"
	"github.com/keyglee/closeio-golang-client/closeio/opportunity"
	"github.com/keyglee/closeio-golang-client/closeio/task"
)

type LeadResponse struct {
	StatusID       string                            `json:"status_id"`
	StatusLabel    string                            `json:"status_label"`
	Tasks          []task.TaskResponse               `json:"tasks"`
	DisplayName    string                            `json:"display_name"`
	Addresses      []Address                         `json:"addresses"`
	Name           string                            `json:"name"`
	Contacts       []contact.ContactResponse         `json:"contacts"`
	CustomField    []customfield.Custom_Field        `json:"custom_field"`
	DateUpdated    string                            `json:"date_updated"`
	HtmlUrl        string                            `json:"html_url"`
	CreatedBy      string                            `json:"created_by"`
	OrganizationID string                            `json:"organization_id"`
	Url            string                            `json:"url"`
	Opportunities  []opportunity.OpportunityResponse `json:"opportunities"`
	UpdatedBy      string                            `json:"updated_by"`
	DateCreated    string                            `json:"date_created"`
	ID             string                            `json:"id"`
	Description    string                            `json:"description"`
}
