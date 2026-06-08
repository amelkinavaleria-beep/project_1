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

type EventController struct {
	service app.EventService
}

func NewEventController(s app.EventService) EventController {
	return EventController{service: s}
}

func (c EventController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		if dev.Category != domain.Actuator {
    	BadRequest(w, errors.New("only actuators can send events"))
    	return
		}

		e, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			BadRequest(w, err)
			return
		}

		e.DeviceId = dev.Id
		e.RoomId = *dev.RoomId

		ev, err := c.service.Save(e)
		if err != nil {
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.DomainToDto(ev))
	}
}

func (c EventController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		events, err := c.service.GetForDevice(dev.Id, *dev.RoomId, page)
		if err != nil {
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.DomainToDtoCollection(events))
	}
}
