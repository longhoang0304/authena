package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/longhoang0304/authena-phnx/internal/controllers"
	nr "github.com/longhoang0304/authena-phnx/libs/named_route"
)

func NewPingPongRoutes(pingPongController controllers.IPingPongController) nr.NamedRoute {
	r := chi.NewRouter()
	r.Get("/", pingPongController.Pong)

	return nr.NamedRoute{
		Name:  "/ping",
		Route: r,
	}
}
