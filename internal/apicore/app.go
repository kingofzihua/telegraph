package apicore

import (
	"github.com/go-ostrich/pkg/app"

	"github.com/kingofzihua/telegraph/internal/apicore/data"

	"github.com/kingofzihua/telegraph/internal/apicore/config"
)

const commandDesc = `telegraph apicore`

func NewApp(basename string) *app.App {
	opts := config.NewOptions()
	application := app.NewApp("telegraph apicore",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(cfg *config.Config) app.RunFunc {
	return func(basename string) error {
		// init db client
		client, err := cfg.MySQL.NewClient()
		if err != nil {
			panic(err)
		}
		data.SetDefaultDBClient(client)

		// Run runs the specified GrpcServer.
		srv, err := createGrpcServer(cfg)
		if err != nil {
			return err
		}

		// This should never exit.
		return srv.PrepareRun().Run()
	}
}
