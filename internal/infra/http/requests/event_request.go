package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventRequest struct {
    Action string `json:"action" validate:"required"`
}

func (r EventRequest) ToDomainModel() (interface{}, error) {
    return domain.Event{Action: r.Action}, nil
}
