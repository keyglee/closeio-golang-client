package note

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (n Note) Update(BaseURL string) (*http.Request, error) {
	updatedNoteBuffer, err := json.Marshal(n)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, BaseURL+"/note/"+n.ID, bytes.NewBuffer(updatedNoteBuffer))

}
