package main

import (
	"context"
	"github.com/longhoang0304/authena-phnx/libs/shutdown"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(shutdown.NewCtx())
	app, err := InitAuthenaApp()

	if err != nil {
		log.Error().Msgf("%v", err)
		return
	}

	eg.Go(func() error {
		return app.Start(ctx)
	})

	defer func() {
		_ = app.Stop(context.Background())
		log.Info().Msg("The authena is shutting down")
	}()

	if err := eg.Wait(); err != nil {
		log.Error().Msgf("%v", err)
	}
}
