package authena

import (
	"github.com/go-chi/chi/v5"
	"github.com/longhoang0304/authena-phnx/libs/named_route"
	"net/http"
)

func NewRouter(routes []named_route.NamedRoute) (http.Handler, error) {
	mainRouter := chi.NewRouter()

	for _, route := range routes {
		mainRouter.Mount(route.Name, route.Route)
	}

	return mainRouter, nil
}
