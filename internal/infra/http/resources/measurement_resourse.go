package resources

import (
	"time"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type MeasurementDto struct {
	Id          uint64    `json:"id"`
	DeviceId    uint64    `json:"deviceId"`
	RoomId      uint64    `json:"roomId"`
	Value       uint64    `json:"value"`
	CreatedDate time.Time `json:"createdDate"`
}

func (d MeasurementDto) DomainToDto(m domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
	}
}

func (d MeasurementDto) DomainToDtoCollection(ms domain.Measurements) []MeasurementDto {
	res := make([]MeasurementDto, len(ms.Items))
	for i, v := range ms.Items {
		res[i] = d.DomainToDto(v)
	}
	return res
}