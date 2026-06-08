package app

import (
    "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
    "github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type EventService interface {
    Save(e domain.Event) (domain.Event, error)
    GetForDevice(deviceId, roomId uint64, page int) (domain.Events, error)
}

type eventService struct {
    repo database.EventRepository
}

func NewEventService(r database.EventRepository) EventService {
    return eventService{repo: r}
}

func (s eventService) Save(e domain.Event) (domain.Event, error) {
    return s.repo.Save(e)
}

func (s eventService) GetForDevice(deviceId, roomId uint64, page int) (domain.Events, error) {
    pagination := domain.Pagination{Page: uint64(page), CountPerPage: 20}
    return s.repo.FindList(pagination, deviceId, roomId)
}
