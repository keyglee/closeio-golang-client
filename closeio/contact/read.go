package contact

import (
	"bytes"
	"net/http"
)

func (c Contact) Read(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodGet, BaseURL+"/contact/"+c.ID, bytes.NewBuffer(nil))

}
