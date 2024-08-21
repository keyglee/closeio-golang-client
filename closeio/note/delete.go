package note

import (
	"bytes"
	"net/http"
)

func (n Note) Delete(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodDelete, BaseURL+BasePath+n.ID, bytes.NewBuffer(nil))

}
