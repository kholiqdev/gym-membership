//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gym/cmd/interface/handler"
	"gym/cmd/interface/handler/health"
	"gym/cmd/interface/handler/member"
	"gym/config"
	"gym/infrastructure/cached"
	"gym/infrastructure/database"
	"gym/internal/protocol/http"
	"gym/internal/protocol/http/router"
	"gym/pkg/auth"
)

func InitHttpProtocol(mode string) (*http.HttpImpl, error) {
	panic(wire.Build(
		config.NewConfig,
		// Wiring for jwt
		wire.NewSet(
			wire.Bind(new(auth.JwtToken), new(*auth.JwtTokenImpl)),
			auth.NewJwtToken,
		),
		// Wiring for data storage
		wire.NewSet(
			database.NewDatabaseClient,
			cached.NewRedisClient,
		),
		// Wiring for http protocol
		wire.NewSet(
			http.NewHttpProtocol,
		),
		// Wiring for http router
		wire.NewSet(
			router.NewHttpRoute,
		),
		// Wiring for http handler
		wire.NewSet(
			handler.NewHttpHandler,
		),
		member.ProviderSet,
		health.ProviderSet,
	))
}
