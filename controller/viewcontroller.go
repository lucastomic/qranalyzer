package controller

import (
	"net/http"

	"github.com/lucastomic/qranalyzer/service"
	"github.com/lucastomic/qranalyzer/view"
)

type VisitorViewController struct {
	visitorService service.VisitorService
}

func NewVisitorViewController(visitorService service.VisitorService) *VisitorViewController {
	return &VisitorViewController{visitorService: visitorService}
}

func (v *VisitorViewController) Visitors(w http.ResponseWriter, r *http.Request) {
	visitors, err := v.visitorService.FindAll()
	if err != nil {
		// handle error
		return
	}

	component := view.Layout(view.Visitors(visitors))
	component.Render(r.Context(), w)
}
