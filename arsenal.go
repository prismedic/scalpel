package arsenal

import (
	"go.uber.org/fx"

	"github.com/prismedic/scalpel/httpfx"
	"github.com/prismedic/scalpel/infofx"
	"github.com/prismedic/scalpel/loggerfx"
	"github.com/prismedic/scalpel/metricsfx"
	"github.com/prismedic/scalpel/routerfx"
)

var Module = fx.Options(
	httpfx.Module,
	infofx.Module,
	loggerfx.Module,
	metricsfx.Module,
	routerfx.Module,
)
