package named_route

import (
	"net/http"
)

type NamedRoute struct {
	Name  string
	Route http.Handler
}
