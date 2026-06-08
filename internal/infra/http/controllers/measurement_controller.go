package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type MeasurementController struct {
	mService app.MeasurementService
}

func NewMeasurementController(ms app.MeasurementService) MeasurementController {
	return MeasurementController{mService: ms}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if dev.Category != domain.Sensor {
			BadRequest(w, errors.New("only sensors can send measurements"))
			return
		}

		m, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			BadRequest(w, err)
			return
		}

		m.DeviceId = dev.Id
		m.RoomId = *dev.RoomId

		m, err = c.mService.Save(m)
		if err != nil {
			InternalServerError(w, err)
			return
		}
		Success(w, resources.MeasurementDto{}.DomainToDto(m))
	}
}

func (c MeasurementController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		period := r.URL.Query().Get("period")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		measurements, err := c.mService.GetForDevice(dev.Id, *dev.RoomId, period, page)
		if err != nil {
			InternalServerError(w, err)
			return
		}
		Success(w, resources.MeasurementDto{}.DomainToDtoCollection(measurements))
	}
}
