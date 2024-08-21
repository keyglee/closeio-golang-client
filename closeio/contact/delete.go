package contact

import (
	"bytes"
	"net/http"
)

func (c Contact) Delete(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodDelete, BaseURL+BasePath+c.GetID(), bytes.NewBuffer(nil))

}
