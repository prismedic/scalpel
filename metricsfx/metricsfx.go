package metricsfx

import (
	"go.uber.org/fx"

	"github.com/prismedic/arsenal/routerfx"
)

var Module = fx.Module("metrics",
	fx.Provide(routerfx.AsHandlerRoute(NewPrometheusHandler)),
)
