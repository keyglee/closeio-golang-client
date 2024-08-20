package note

type NoteResponse struct {
	Attachments    []NoteAttachmentResponse `json:"attachments"`
	Type           string                   `json:"_type"`
	NoteHtml       string                   `json:"note_html"`
	Note           string                   `json:"note"`
	UserId         string                   `json:"user_id"`
	UserName       string                   `json:"user_name"`
	UpdatedBy      string                   `json:"updated_by"`
	UpdatedByName  string                   `json:"updated_by_name"`
	DateUpdated    string                   `json:"date_updated"`
	CreatedBy      string                   `json:"created_by"`
	CreatedByName  string                   `json:"created_by_name"`
	OrganizationId string                   `json:"organization_id"`
	ContactId      string                   `json:"contact_id"`
	DateCreated    string                   `json:"date_created"`
	ID             string                   `json:"id"`
	LeadId         string                   `json:"lead_id"`
}

type NoteAttachmentResponse struct {
	ContentType  string `json:"content_type"`
	Filename     string `json:"filename"`
	Size         int    `json:"size"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnail_url"`
}
