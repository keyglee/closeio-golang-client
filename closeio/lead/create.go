package lead

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (l Lead) Create(BaseURL string) (*http.Request, error) {
	newLeadBuffer, err := json.Marshal(l)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, BaseURL+"/lead/", bytes.NewBuffer(newLeadBuffer))

}
