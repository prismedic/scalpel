package routerfx

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("router",
	fx.Provide(New),
	fx.Provide(AsControllerRoute(NewSwaggerController)),
)

type Config struct {
	CorsAllowedOrigins []string `mapstructure:"cors_allowed_origins" yaml:"cors_allowed_origins"`
}

type Params struct {
	fx.In
	Config           *Config
	Logger           *zap.SugaredLogger `optional:"true"`
	ControllerRoutes []ControllerRoute  `group:"controllerRoutes"`
	HandlerRoutes    []HandlerRoute     `group:"handlerRoutes"`
	Middlewares      []gin.HandlerFunc  `group:"middlewares"`
	NoRouteHandler   gin.HandlerFunc    `optional:"true" name:"noRouteHandler"`
}

type Result struct {
	fx.Out
	Router http.Handler
}

func New(p Params) Result {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	if p.Logger != nil {
		router.Use(ginzap.Ginzap(p.Logger.Desugar(), time.RFC3339, true))
	}
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = p.Config.CorsAllowedOrigins
	router.Use(cors.New(corsConfig))

	for _, middleware := range p.Middlewares {
		router.Use(middleware)
	}

	for _, route := range p.ControllerRoutes {
		if p.Logger != nil {
			p.Logger.Infow("registering controller route", "pattern", route.RoutePattern())
		}
		route.RegisterControllerRoutes(
			router.Group(route.RoutePattern()),
		)
	}

	for _, route := range p.HandlerRoutes {
		if p.Logger != nil {
			p.Logger.Infow("registering handler route", "pattern", route.RoutePattern())
		}
		router.Any(route.RoutePattern(), route.Handler())
	}

	if p.NoRouteHandler != nil {
		if p.Logger != nil {
			p.Logger.Info("registering no route handler")
		}
		router.NoRoute(p.NoRouteHandler)
	}

	return Result{
		Router: router,
	}
}

func (r *Result) GetHttpRouter() http.Handler {
	return r.Router
}

type ControllerRoute interface {
	RegisterControllerRoutes(rg *gin.RouterGroup)
	RoutePattern() string
}

func AsControllerRoute(controller any) any {
	return fx.Annotate(
		controller,
		fx.As(new(ControllerRoute)),
		fx.ResultTags(`group:"controllerRoutes"`),
	)
}

type HandlerRoute interface {
	Handler() gin.HandlerFunc
	RoutePattern() string
}

func AsHandlerRoute(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(HandlerRoute)),
		fx.ResultTags(`group:"handlerRoutes"`),
	)
}
