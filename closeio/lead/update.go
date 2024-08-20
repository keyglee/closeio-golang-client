package lead

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (l Lead) Update(BaseURL string) (*http.Request, error) {
	updatedLead, err := json.Marshal(l)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, BaseURL+"/lead/"+l.ID, bytes.NewBuffer(updatedLead))

}
