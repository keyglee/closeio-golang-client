package contact

type Contact struct {
	Name   string  `json:"name,omitempty"`
	Title  string  `json:"title,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Phones []Email `json:"phones,omitempty"`
}

type Email struct {
	Type  string `json:"type,omitempty"`
	Email string `json:"email,omitempty"`
}

type Phone struct {
	Type  string `json:"type,omitempty"`
	Phone string `json:"phone,omitempty"`
}
