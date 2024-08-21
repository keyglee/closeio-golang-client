package task

import (
	"bytes"
	"net/http"
)

func (t Task) Read(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodGet, BaseURL+BasePath+t.GetID(), bytes.NewBuffer(nil))

}
