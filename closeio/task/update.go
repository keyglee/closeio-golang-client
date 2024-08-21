package task

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (t Task) Update(BaseURL string) (*http.Request, error) {
	updatedTaskBuffer, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, BaseURL+BasePath+t.GetID(), bytes.NewBuffer(updatedTaskBuffer))

}
