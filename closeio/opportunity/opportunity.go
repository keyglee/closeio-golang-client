package opportunity

import customfield "github.com/keyglee/closeio-golang-client/closeio/customField"

type Opportunity struct {
	ID           string                    `json:"-"`
	Note         string                    `json:"note"`
	Confidence   int                       `json:"confidence"`
	Lead_id      string                    `json:"lead_id"`
	Status_id    string                    `json:"status_id"`
	Value        int                       `json:"value"`
	Value_period string                    `json:"value_period"`
	Custom       []customfield.CustomField `json:"custom"`
}
