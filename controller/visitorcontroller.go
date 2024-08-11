package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lucastomic/qranalyzer/domain"
	"github.com/lucastomic/qranalyzer/service"
)

type VisitorController struct {
	visitorService service.VisitorService
}

func NewVisitorController(visitorService service.VisitorService) *VisitorController {
	return &VisitorController{visitorService}
}

func (c *VisitorController) Create(w http.ResponseWriter, r *http.Request) {
	var visitor domain.Visitor
	err := json.NewDecoder(r.Body).Decode(&visitor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.visitorService.Create(&visitor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(visitor)
}
