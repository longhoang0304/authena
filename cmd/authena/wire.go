//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/longhoang0304/authena-phnx/internal/apps/authena"
)

func InitAuthenaApp() (*authena.AuthenaApp, error) {
	wire.Build(authena.AuthenaAppGraphSet)
	return &authena.AuthenaApp{}, nil
}
