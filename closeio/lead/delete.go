package lead

import (
	"bytes"
	"net/http"
)

func (l Lead) Delete(BaseURL string) (*http.Request, error) {

	return http.NewRequest(http.MethodDelete, BaseURL+BasePath+l.GetID(), bytes.NewBuffer(nil))

}
