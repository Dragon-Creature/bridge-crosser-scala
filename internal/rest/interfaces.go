package rest

import (
	"net/http"
)

//go:generate moq -out mock.go . Interfaces
type Interfaces interface {
	PostCalculateCrossing(w http.ResponseWriter, r *http.Request)
}
