package controller

import (
	"net/http"

	"github.com/lucastomic/qranalyzer/view"
)

type MiddlewareController struct {
}

func NewMiddlewareController() *MiddlewareController {
	return &MiddlewareController{}
}

func (m *MiddlewareController) LoggerMiddleware(w http.ResponseWriter, r *http.Request) {
	component := view.LoggerMiddleware()
	component.Render(r.Context(), w)
}
