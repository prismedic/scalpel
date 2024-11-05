package routerfx

import (
	"net/http"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DocsContent struct {
	Title       string `json:"title"`
	SpecContent string `json:"spec_content"`
}

type ScalarParams struct {
	fx.In
	DocsContent *DocsContent `optional:"true"`
	Logger      *zap.SugaredLogger
}

type ScalarHandler struct {
	DocsHTML string
}

func NewScalarHandler(p ScalarParams) (*ScalarHandler, error) {
	if p.DocsContent == nil {
		p.Logger.Warn("No docs content provided, docs will not be available")
		return &ScalarHandler{}, nil
	}
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecContent: p.DocsContent.SpecContent,
		CustomOptions: scalar.CustomOptions{
			PageTitle: p.DocsContent.Title,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ScalarHandler{
		DocsHTML: htmlContent,
	}, nil
}

func (h *ScalarHandler) Docs(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, h.DocsHTML)
}

func (h *ScalarHandler) RegisterControllerRoutes(rg *gin.RouterGroup) {
	// enable docs only if docs is provided
	if h.DocsHTML != "" {
		rg.GET("", h.Docs)
	}
}

func (h *ScalarHandler) RoutePattern() string {
	// TODO: handle future versions of API docs (e.g. v2)
	return "/v1/docs"
}
