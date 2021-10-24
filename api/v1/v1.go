package v1

import (
	"khidr/service-b/service"
	"net/http"

	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
)

type V1 struct {
	svc *service.Service
}

func NewV1(svc *service.Service) *V1 {
	return &V1{
		svc: svc,
	}
}

func (v *V1) GetDailyManifest(w http.ResponseWriter, r *http.Request) {
	cargosResponse, err := v.svc.GetDailyManifest()
	if err != nil {
		handleError(w, r, err, "", http.StatusInternalServerError)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, cargosResponse)
}

func handleError(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	if message == "" {
		message = "Internal Error"
	}
	errors := Error{Errors: []string{message}}
	log.Error(errors)
	render.Status(r, status)
	render.JSON(w, r, errors)
}
