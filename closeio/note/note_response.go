package note

type NoteResponse struct {
	Attachments     []NoteAttachmentResponse `json:"attachments"`
	Type            string                   `json:"_type"`
	NoteHtml        string                   `json:"note_html"`
	Note            string                   `json:"note"`
	User_id         string                   `json:"user_id"`
	User_name       string                   `json:"user_name"`
	Updated_by      string                   `json:"updated_by"`
	Updated_by_name string                   `json:"updated_by_name"`
	Date_updated    string                   `json:"date_updated"`
	Created_by      string                   `json:"created_by"`
	Created_by_name string                   `json:"created_by_name"`
	Organization_id string                   `json:"organization_id"`
	Contact_id      string                   `json:"contact_id"`
	Date_created    string                   `json:"date_created"`
	Id              string                   `json:"id"`
	Lead_id         string                   `json:"lead_id"`
}

type NoteAttachmentResponse struct {
	Content_type  string `json:"content_type"`
	Filename      string `json:"filename"`
	Size          int    `json:"size"`
	Url           string `json:"url"`
	Thumbnail_url string `json:"thumbnail_url"`
}
