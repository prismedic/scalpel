package infofx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

type HealthResponse struct {
	Status string `json:"status"`
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// getHealth godoc
//	@Summary		Get health status
//	@Description	Get health status of the service
//	@Produce		json
//	@Success		200	{object}	HealthResponse
//	@Router			/healthz [get]
func (hc *HealthController) getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, &HealthResponse{Status: "OK"})
}

func (hc *HealthController) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.GET("/", hc.getHealth)
}

func (hc *HealthController) RoutePattern() string {
	return "/healthz"
}
