package closeio

import "net/http"

type CloseModel interface {
	GetID() string
	Create(BaseURL string) (*http.Request, error)
	Update(BaseURL string) (*http.Request, error)
	Read(BaseURL string) (*http.Request, error)
	Delete(BaseURL string) (*http.Request, error)
}
