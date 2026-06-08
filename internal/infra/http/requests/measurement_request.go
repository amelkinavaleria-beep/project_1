package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementRequest struct {
	Value uint64 `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		Value: r.Value,
	}, nil
}