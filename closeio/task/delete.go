package task

import (
	"bytes"
	"net/http"
)

func (t Task) Delete(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodDelete, BaseURL+BasePath+t.GetID(), bytes.NewBuffer(nil))

}
