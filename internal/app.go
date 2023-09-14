package internal

import (
	"github.com/go-ostrich/pkg/app"

	"github.com/kingofzihua/telegraph/internal/config"
)

const commandDesc = `Imitate telegra.ph website`

func NewApp(basename string) *app.App {
	opts := config.NewOptions()
	application := app.NewApp("telegra.ph",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(cfg *config.Config) app.RunFunc {
	return func(basename string) error {
		// Run runs the specified APIServer. This should never exit.
		srv, err := createAPIServer(cfg)
		if err != nil {
			return err
		}

		return srv.PrepareRun().Run()
	}
}
