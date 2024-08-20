package closeio

type CloseLeadWebhookEvent struct {
	Id             string                 `json:"id"`
	DateCreated    string                 `json:"date_created"`
	DateUpdated    string                 `json:"date_updated"`
	OrganizationId string                 `json:"organization_id"`
	UserId         string                 `json:"user_id"`
	RequestId      string                 `json:"request_id"`
	ObjectType     string                 `json:"object_type"`
	ObjectId       string                 `json:"object_id"`
	LeadId         string                 `json:"lead_id"`
	Action         string                 `json:"action"`
	ChangedFields  []string               `json:"changed_fields"`
	Data           map[string]interface{} `json:"data"`
	PreviousData   map[string]interface{} `json:"previous_data"`
}
