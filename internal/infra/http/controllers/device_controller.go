package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type DeviceController struct {
	deviceService app.DeviceService
}

func NewDeviceController(ds app.DeviceService) DeviceController {
	return DeviceController{
		deviceService: ds,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)

		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		dev.OrganizationId = org.Id
		dev, err = c.deviceService.Save(dev)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

func (c DeviceController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)

		devices, err := c.deviceService.FindList(org.Id)
		if err != nil {
			log.Printf("DeviceController.FindList: %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDtoCollection(devices))
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newDevData, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Update(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		dev.InventoryNumber = newDevData.InventoryNumber
		dev.SerialNumber = newDevData.SerialNumber
		dev.Characteristics = newDevData.Characteristics
		dev.RoomId = newDevData.RoomId
		dev.Units = newDevData.Units
		dev.PowerConsumption = newDevData.PowerConsumption

		dev, err = c.deviceService.Update(dev)
		if err != nil {
			log.Printf("DeviceController.Update(c.deviceService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		err := c.deviceService.Delete(dev.Id)
		if err != nil {
			log.Printf("DeviceController.Delete: %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}