package metricsfx

import (
	"github.com/gin-gonic/gin"
	"github.com/prismedic/scalpel/routerfx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHandler struct{}

func NewPrometheusHandler() *PrometheusHandler {
	return &PrometheusHandler{}
}

func (ph *PrometheusHandler) Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func (ph *PrometheusHandler) RoutePattern() string {
	return "/metrics"
}

var _ routerfx.HandlerRoute = (*PrometheusHandler)(nil)
