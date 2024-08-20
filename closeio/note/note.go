package note

type Note struct {
	Note_html   string           `json:"note_html,omitempty"`
	Lead_id     string           `json:"lead_id,omitempty"`
	Attachments []NoteAttachment `json:"attachments,omitempty"`
}

type NoteAttachment struct {
	Content_type string `json:"content_type,omitempty"`
	Filename     string `json:"filename,omitempty"`
	Url          string `json:"url,omitempty"`
}
