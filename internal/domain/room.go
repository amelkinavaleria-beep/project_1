package domain

import "time"

type Room struct {
	Id             uint64
	OrganizationId uint64
	Name           string
	Description    *string
	CreatedDate    time.Time
	UpdatedDate    time.Time
	DeletedDate    *time.Time
}

type Rooms struct {
	Items []Room
	Total uint64
	Pages uint
}