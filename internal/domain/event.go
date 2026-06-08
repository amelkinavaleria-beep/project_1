package domain

import "time"

type Event struct {
    Id          uint64
    DeviceId    uint64
    RoomId      uint64
    Action      string // "ON" або "OFF"
    CreatedDate time.Time
    UpdatedDate time.Time
    DeletedDate *time.Time
}

type Events struct {
    Items []Event
    Total uint64
    Pages uint
}
