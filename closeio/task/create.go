package task

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (t Task) Create(BaseURL string) (*http.Request, error) {
	newTaskBuffer, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, BaseURL+BasePath, bytes.NewBuffer(newTaskBuffer))

}
