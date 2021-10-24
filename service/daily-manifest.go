package service

import "math"

type CargosResponse struct {
	Cargos []Cargo `json:"cargos"`
}

type Cargo struct {
	OrderIds    []int   `json:"order_ids"`
	TotalWeight float64 `json:"total_weight"`
}

func (s *Service) GetDailyManifest() (*CargosResponse, error) {
	cargos := make([]Cargo, 0)
	collections := []string{"Cameroon", "Ethiopia", "Morocco", "Mozambique", "Uganda"}
	for _, collection := range collections {
		err := s.getCargosByCollection(collection, &cargos)
		if err != nil {
			return nil, err
		}
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
