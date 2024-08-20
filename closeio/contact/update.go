package contact

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c Contact) Update(BaseURL string) (*http.Request, error) {
	updatedContactBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, BaseURL+"/contact/"+c.ID, bytes.NewBuffer(updatedContactBuffer))

}
