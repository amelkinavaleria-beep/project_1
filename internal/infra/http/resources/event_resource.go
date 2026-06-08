package resources

import (
    "time"
    "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type EventDto struct {
    Id          uint64    `json:"id"`
    DeviceId    uint64    `json:"deviceId"`
    RoomId      uint64    `json:"roomId"`
    Action      string    `json:"action"`
    CreatedDate time.Time `json:"createdDate"`
}

func (d EventDto) DomainToDto(e domain.Event) EventDto {
    return EventDto{
        Id:          e.Id,
        DeviceId:    e.DeviceId,
        RoomId:      e.RoomId,
        Action:      e.Action,
        CreatedDate: e.CreatedDate,
    }
}

func (d EventDto) DomainToDtoCollection(ev domain.Events) []EventDto {
    res := make([]EventDto, len(ev.Items))
    for i, v := range ev.Items {
        res[i] = d.DomainToDto(v)
    }
    return res
}
