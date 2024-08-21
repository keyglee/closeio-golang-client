package contact

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c Contact) Create(BaseURL string) (*http.Request, error) {
	newContactBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, BaseURL+BasePath, bytes.NewBuffer(newContactBuffer))

}
