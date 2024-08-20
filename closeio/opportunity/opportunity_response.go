package opportunity

type OpportunityResponse struct {
	StatusId       string  `json:"status_id"`
	StatusLabel    string  `json:"status_label"`
	StatusType     string  `json:"status_type"`
	PipelineId     string  `json:"pipeline_id"`
	PipelineName   string  `json:"pipeline_name"`
	DateWon        string  `json:"date_won"`
	Confidence     int     `json:"confidence"`
	UserId         string  `json:"user_id"`
	ContactId      string  `json:"contact_id"`
	UpdatedBy      string  `json:"updated_by"`
	DateUpdated    string  `json:"date_updated"`
	ValuePeriod    string  `json:"value_period"`
	CreatedBy      string  `json:"created_by"`
	Note           string  `json:"note"`
	Value          float64 `json:"value"`
	ValueFormatted string  `json:"value_formatted"`
	ValueCurrency  string  `json:"value_currency"`
	LeadName       string  `json:"lead_name"`
	OrganizationNd string  `json:"organization_id"`
	DateCreated    string  `json:"date_created"`
	UserName       string  `json:"user_name"`
	ID             string  `json:"id"`
	LeadId         string  `json:"lead_id"`
}
