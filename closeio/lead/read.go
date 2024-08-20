package lead

import (
	"bytes"
	"net/http"
)

func (l Lead) Read(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodGet, BaseURL+"/lead/"+l.ID, bytes.NewBuffer(nil))

}
