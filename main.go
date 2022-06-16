package main

import (
	"context"

	"github.com/jgcaceres97/go-inventory-api/src/database"
	"github.com/jgcaceres97/go-inventory-api/src/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")

				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
