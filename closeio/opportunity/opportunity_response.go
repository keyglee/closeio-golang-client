package opportunity

type OpportunityResponse struct {
	Status_id       string  `json:"status_id"`
	Status_label    string  `json:"status_label"`
	Status_type     string  `json:"status_type"`
	Pipeline_id     string  `json:"pipeline_id"`
	Pipeline_name   string  `json:"pipeline_name"`
	Date_won        string  `json:"date_won"`
	Confidence      int     `json:"confidence"`
	User_id         string  `json:"user_id"`
	Contact_id      string  `json:"contact_id"`
	Updated_by      string  `json:"updated_by"`
	Date_updated    string  `json:"date_updated"`
	Value_period    string  `json:"value_period"`
	Created_by      string  `json:"created_by"`
	Note            string  `json:"note"`
	Value           float64 `json:"value"`
	Value_formatted string  `json:"value_formatted"`
	Value_currency  string  `json:"value_currency"`
	Lead_name       string  `json:"lead_name"`
	Organization_id string  `json:"organization_id"`
	Date_created    string  `json:"date_created"`
	User_name       string  `json:"user_name"`
	Id              string  `json:"id"`
	Lead_id         string  `json:"lead_id"`
}
