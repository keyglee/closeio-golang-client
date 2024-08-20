package note

import (
	"bytes"
	"net/http"
)

func (n Note) Read(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodGet, BaseURL+"/note/"+n.ID, bytes.NewBuffer(nil))

}
