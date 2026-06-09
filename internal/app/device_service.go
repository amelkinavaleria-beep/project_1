package app

import (
	"errors"
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (interface{}, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
	FindList(orgId uint64, page int) (domain.Devices, error)
}

type deviceService struct {
	deviceRepo database.DeviceRepository
}

func NewDeviceService(dr database.DeviceRepository) DeviceService {
	return deviceService{
		deviceRepo: dr,
	}
}

func (s deviceService) Save(d domain.Device) (domain.Device, error) {
	if err := s.validateDevice(d); err != nil {
		return domain.Device{}, err
	}

	dev, err := s.deviceRepo.Save(d)
	if err != nil {
		log.Printf("deviceService.Save: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	dev, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.Find: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) Update(d domain.Device) (domain.Device, error) {
	if err := s.validateDevice(d); err != nil {
		return domain.Device{}, err
	}

	dev, err := s.deviceRepo.Update(d)
	if err != nil {
		log.Printf("deviceService.Update: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("deviceService.Delete: %s", err)
		return err
	}
	return nil
}

func (s deviceService) FindList(orgId uint64, page int) (domain.Devices, error) {
	pagination := domain.Pagination{
		Page:         uint64(page),
		CountPerPage: 20,
	}
	return s.deviceRepo.FindList(pagination, orgId)
}

func (s deviceService) validateDevice(d domain.Device) error {
	if d.Category == domain.Sensor && d.Units == "" {
		return errors.New("для вимірювальних пристроїв (SENSOR) має бути заповнене поле Units")
	}
	if d.Category == domain.Actuator && d.PowerConsumption == 0 {
		return errors.New("для виконавчих пристроїв (ACTUATOR) має бути заповнене поле PowerConsumption")
	}
	return nil
}