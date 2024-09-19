package app

import (
	"github.com/prismedic/scalpel"
	"go.uber.org/fx"

	"examples/httpServer/internal/pkg/config"
)

func New() *fx.App {
	app := fx.New(
		arsenal.Module,
		config.Module,
	)
	return app
}
