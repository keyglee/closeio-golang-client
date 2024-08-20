package note

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (n Note) Create(BaseURL string) (*http.Request, error) {
	newNoteBuffer, err := json.Marshal(n)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, BaseURL+"/note/", bytes.NewBuffer(newNoteBuffer))

}
