package authena

import (
	"context"
	"github.com/longhoang0304/authena-phnx/libs/shutdown"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type AuthenaApp struct {
	httpServer *HTTPServer
}

func (self *AuthenaApp) Start(ctx context.Context) error {
	eg, childCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return shutdown.BlockListen(childCtx, func() error {
			log.Info().Msgf("Api's listening on %s", self.httpServer.Addr)
			return self.httpServer.ListenAndServe()
		})
	})

	return eg.Wait()
}

func (self *AuthenaApp) Stop(ctx context.Context) error {
	return self.httpServer.Shutdown(ctx)
}

func NewAuthenaApp(httpServer *HTTPServer) *AuthenaApp {
	return &AuthenaApp{
		httpServer,
	}
}
