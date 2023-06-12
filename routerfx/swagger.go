package routerfx

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct{}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{}
}

func (sc *SwaggerController) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (sc *SwaggerController) RoutePattern() string {
	return "/docs"
}
