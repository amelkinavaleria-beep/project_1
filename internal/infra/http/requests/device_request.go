package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	RoomId           *uint64                `json:"roomId"`
	GUID             string                 `json:"guid" validate:"required"`
	InventoryNumber  string                 `json:"inventoryNumber" validate:"required"`
	SerialNumber     string                 `json:"serialNumber" validate:"required"`
	Characteristics  string                 `json:"characteristics"`
	Category         domain.DeviceCategory  `json:"category" validate:"required"`
	Units            string                 `json:"units"`
	PowerConsumption float64                `json:"powerConsumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		RoomId:           r.RoomId,
		GUID:             r.GUID,
		InventoryNumber:  r.InventoryNumber,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         r.Category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}