package api

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/rest"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := &service.Service{}
	re := &rest.Rest{
		Service: s,
	}

	r.Post("/api/calculate_crossing", re.PostCalculateCrossing)
	fileServer := http.FileServer(http.Dir("./build/"))
	r.Handle("/*", fileServer)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
