package authena

import (
	"github.com/google/wire"
	"github.com/longhoang0304/authena-phnx/internal/controllers"
	"github.com/longhoang0304/authena-phnx/internal/routes"
	"github.com/longhoang0304/authena-phnx/libs/named_route"
)

func NewRoutes() []named_route.NamedRoute {
	pingPongController := controllers.NewPingPongController()
	pingPongRoute := routes.NewPingPongRoutes(pingPongController)

	return []named_route.NamedRoute{
		pingPongRoute,
	}
}

var InitRouter = wire.NewSet(
	NewRoutes,
	NewRouter,
)

var AuthenaAppGraphSet = wire.NewSet(
	InitRouter,
	NewHttpServer,
	NewAuthenaApp,
)
