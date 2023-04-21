package api

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/calculate_crossing", rest.PostCalculateCrossing)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
