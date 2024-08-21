package contact

import (
	"bytes"
	"net/http"
)

func (c Contact) Read(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodGet, BaseURL+BasePath+c.GetID(), bytes.NewBuffer(nil))

}
