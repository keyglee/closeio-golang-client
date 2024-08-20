package task

type TaskResponse struct {
	Type           string `json:"_type"`
	AssignedTo     string `json:"assigned_to"`
	AssignedToName string `json:"assigned_to_name"`
	ContactId      string `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	CreatedBy      string `json:"created_by"`
	CreatedByName  string `json:"created_by_name"`
	Date           string `json:"date"`
	DateCreated    string `json:"date_created"`
	DateUpdated    string `json:"date_updated"`
	ID             string `json:"id"`
	IsComplete     string `json:"is_complete"`
	IsDateless     string `json:"is_dateless"`
	LeadId         string `json:"lead_id"`
	LeadName       string `json:"lead_name"`
	ObjectId       string `json:"object_id"`
	ObjectType     string `json:"object_type"`
	OrganizationId string `json:"organization_id"`
	Text           string `json:"text"`
	UpdatedBy      string `json:"updated_by"`
	UpdatedByName  string `json:"updated_by_name"`
}
