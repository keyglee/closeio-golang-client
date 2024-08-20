package task

type Task struct {
	ID         string `json:"-"`
	Type       string `json:"_type,omitempty"`
	LeadId     string `json:"lead_id,omitempty"`
	AssignedTo string `json:"assigned_to,omitempty"`
	Text       string `json:"text,omitempty"`
	Date       string `json:"date,omitempty"`
	IsComplete bool   `json:"is_complete,omitempty"`
}
