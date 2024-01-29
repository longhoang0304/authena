package authena

import (
	"net/http"
)

func NewHttpServer(router http.Handler) (*HTTPServer, error) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	return &HTTPServer{Server: server}, nil
}

type HTTPServer struct {
	*http.Server
}
