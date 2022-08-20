package main

import (
	"context"

	"github.com/jgcaceres97/go-inventory-api/src/database"
	"github.com/jgcaceres97/go-inventory-api/src/internal/repository"
	"github.com/jgcaceres97/go-inventory-api/src/internal/service"
	"github.com/jgcaceres97/go-inventory-api/src/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(),
	)

	app.Run()
}
