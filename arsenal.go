package arsenal

import (
	"go.uber.org/fx"

	"github.com/prismedic/arsenal/httpfx"
	"github.com/prismedic/arsenal/infofx"
	"github.com/prismedic/arsenal/loggerfx"
	"github.com/prismedic/arsenal/metricsfx"
	"github.com/prismedic/arsenal/routerfx"
)

var Module = fx.Options(
	httpfx.Module,
	infofx.Module,
	loggerfx.Module,
	metricsfx.Module,
	routerfx.Module,
)
