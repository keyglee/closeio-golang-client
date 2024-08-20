package task

type Task struct {
	Type        string `json:"_type,omitempty"`
	Lead_id     string `json:"lead_id,omitempty"`
	Assigned_to string `json:"assigned_to,omitempty"`
	Text        string `json:"text,omitempty"`
	Date        string `json:"date,omitempty"`
	Is_complete bool   `json:"is_complete,omitempty"`
}
