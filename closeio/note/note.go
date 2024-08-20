package note

type Note struct {
	ID          string           `json:"-"`
	NoteHtml    string           `json:"note_html,omitempty"`
	LeadId      string           `json:"lead_id,omitempty"`
	Attachments []NoteAttachment `json:"attachments,omitempty"`
}

type NoteAttachment struct {
	ContentType string `json:"content_type,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Url         string `json:"url,omitempty"`
}
