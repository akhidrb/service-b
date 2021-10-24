package service

import (
	"math"

	"golang.org/x/sync/errgroup"
)

type CargosResponse struct {
	Cargos []Cargo `json:"cargos"`
}

type Cargo struct {
	OrderIds    []int   `json:"order_ids"`
	TotalWeight float64 `json:"total_weight"`
}

func (s *Service) GetDailyManifest() (*CargosResponse, error) {
	cargos := make([]Cargo, 0)
	var eg errgroup.Group
	collections := []string{"Cameroon", "Ethiopia", "Morocco", "Mozambique", "Uganda"}
	for i := range collections {
		collection := collections[i]
		eg.Go(func() error {
			err := s.getCargosByCollection(collection, &cargos)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	cargosResponse := CargosResponse{
		Cargos: cargos,
	}
	return &cargosResponse, nil
}

func (s *Service) getCargosByCollection(collection string, cargos *[]Cargo) error {
	orders, err := s.MongoConfig.GetDailyManifest(collection)
	if err != nil {
		return err
	}
	for _, order := range orders {
		cargo := Cargo{
			OrderIds:    order.OrderIds,
			TotalWeight: math.Round(order.TotalWeight*100) / 100,
		}
		*cargos = append(*cargos, cargo)
	}
	return nil
}
