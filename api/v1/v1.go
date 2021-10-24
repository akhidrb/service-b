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

type CargosResponse struct {
	Cargos []Cargo `json:"cargos"`
}

type Cargo struct {
	OrderIds    []int   `json:"order_ids"`
	TotalWeight float64 `json:"total_weight"`
}

func (v *V1) GetDailyManifest(w http.ResponseWriter, r *http.Request) {
	orders, err := v.svc.MongoConfig.GetDailyManifest("Cameroon")
	if err != nil {
		handleError(w, r, err, "", http.StatusInternalServerError)
		return
	}
	cargos := make([]Cargo, 0)
	for _, order := range orders {
		cargo := Cargo{
			OrderIds:    order.OrderIds,
			TotalWeight: order.TotalWeight,
		}
		cargos = append(cargos, cargo)
	}
	cargosResponse := CargosResponse{
		Cargos: cargos,
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
