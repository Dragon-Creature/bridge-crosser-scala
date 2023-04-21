package service

import "git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"

//go:generate moq -out mock.go . Interfaces
type Interfaces interface {
	CalculateCrossing(request model.CrossingRequest) model.CrossingResponse
}
