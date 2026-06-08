package app

import (
	"time"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	GetForDevice(deviceId, roomId uint64, period string, page int) (domain.Measurements, error)
}

type measurementService struct {
	mRepo database.MeasurementRepository
}

func NewMeasurementService(mr database.MeasurementRepository) MeasurementService {
	return measurementService{mRepo: mr}
}

func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	return s.mRepo.Save(m)
}

func (s measurementService) GetForDevice(deviceId, roomId uint64, period string, page int) (domain.Measurements, error) {
	now := time.Now()
	var from time.Time

	// Логіка вибору періоду згідно з Етапом 5 [1]
	switch period {
	case "day":
		from = now.AddDate(0, 0, -1)
	case "week":
		from = now.AddDate(0, 0, -7)
	case "month":
		from = now.AddDate(0, -1, 0)
	default:
		from = now.AddDate(0, 0, -1) // за замовчуванням день
	}

	filters := database.MeasurementFilters{
		DeviceId:        deviceId,
		RoomId:          roomId,
		CreatedDateFrom: &from,
		CreatedDateTo:   &now,
	}

	pagination := domain.Pagination{Page: uint64(page), CountPerPage: 20}
	return s.mRepo.FindList(pagination, filters)
}